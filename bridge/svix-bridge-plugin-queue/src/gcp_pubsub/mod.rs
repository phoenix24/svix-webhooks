use crate::error::Error;
use crate::{run_inner, Consumer};
use omniqueue::{
    backends,
    queue::{consumer::DynConsumer, QueueBackend},
};
use serde::Deserialize;
use std::path::PathBuf;
use svix_bridge_types::{
    async_trait, svix::api::Svix, SenderInput, SenderOutputOpts, TransformationConfig,
    TransformerTx,
};

#[derive(Debug, Default, Deserialize)]
pub struct GCPPubSubInputOpts {
    pub subscription_id: String,
    pub credentials_file: Option<PathBuf>,
}

pub struct GCPPubSubConsumerPlugin {
    name: String,
    input_options: GCPPubSubInputOpts,
    svix_client: Svix,
    transformer_tx: Option<TransformerTx>,
    transformation: Option<TransformationConfig>,
}

impl GCPPubSubConsumerPlugin {
    pub fn new(
        name: String,
        input: GCPPubSubInputOpts,
        transformation: Option<TransformationConfig>,
        output: SenderOutputOpts,
    ) -> Self {
        Self {
            name,
            input_options: input,
            svix_client: match output {
                SenderOutputOpts::Svix(output) => {
                    Svix::new(output.token, output.options.map(Into::into))
                }
            },
            transformer_tx: None,
            transformation,
        }
    }
}

#[async_trait]
impl Consumer for GCPPubSubConsumerPlugin {
    fn source(&self) -> &str {
        &self.input_options.subscription_id
    }

    fn system(&self) -> &str {
        "gcp-pubsub"
    }

    fn transformer_tx(&self) -> &Option<TransformerTx> {
        &self.transformer_tx
    }

    fn transformation(&self) -> &Option<TransformationConfig> {
        &self.transformation
    }

    fn svix_client(&self) -> &Svix {
        &self.svix_client
    }

    async fn consumer(&self) -> std::io::Result<DynConsumer> {
        let consumer = backends::gcp_pubsub::GcpPubSubBackend::builder(
            backends::gcp_pubsub::GcpPubSubConfig {
                subscription_id: self.input_options.subscription_id.clone(),
                credentials_file: self.input_options.credentials_file.clone(),
                // Topics are for producers so we don't care
                topic_id: String::new(),
            },
        )
        .make_dynamic()
        .build_consumer()
        .await
        .map_err(Error::from)?;
        Ok(consumer)
    }
}

#[async_trait]
impl SenderInput for GCPPubSubConsumerPlugin {
    fn name(&self) -> &str {
        &self.name
    }
    fn set_transformer(&mut self, tx: Option<TransformerTx>) {
        self.transformer_tx = tx;
    }
    async fn run(&self) -> std::io::Result<()> {
        run_inner(self).await
    }
}

#[derive(Clone, Debug, Deserialize)]
pub struct GCPPubSubOutputOpts {
    pub topic: String,
    pub credentials_file: Option<PathBuf>,
}
