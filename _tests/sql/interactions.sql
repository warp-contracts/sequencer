create table interactions
(
    id                  serial,
    interaction_id      varchar(255)                                      not null,
    interaction         jsonb                                             not null,
    block_height        integer                                           not null,
    block_id            varchar(255)                                      not null,
    contract_id         varchar(255)                                      not null,
    function            varchar(255),
    input               text,
    confirmation_status varchar(255)                                      not null,
    confirming_peer     varchar(255),
    confirmed_at_height bigint,
    confirmations       varchar(255),
    source              varchar(255) default 'arweave'::character varying not null,
    bundler_tx_id       varchar(255) default NULL::character varying,
    interact_write      text[],
    sort_key            varchar(255) default ''::character varying        not null,
    evolve              varchar(64)
);
