create table sequencer
(
    id                      serial,
    original_sig            text                  not null,
    original_owner          text                  not null,
    sequence_block_id       text                  not null,
    sequence_block_height   integer               not null,
    sequence_transaction_id text                  not null,
    bundler_tx_id           text                  not null,
    bundler_response        text                  not null,
    original_address        text                  not null,
    sequence_sort_key       text default ''::text not null,
    sequence_millis         text default ''::text not null
);
