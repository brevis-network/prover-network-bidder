CREATE DATABASE bidder;

CREATE USER bidder;

SET DATABASE TO bidder;

CREATE TABLE IF NOT EXISTS app (
    app_id TEXT NOT NULL, -- hex of vk, bytes32
    img_url TEXT NOT NULL,
    register_status TEXT NOT NULL DEFAULT '',
    register_error TEXT NOT NULL DEFAULT '',
    PRIMARY KEY (app_id)
);

CREATE TABLE IF NOT EXISTS proof_request (
    req_id TEXT NOT NULL, -- keccak256(abi.encodePacked(req.nonce, req.vk, req.publicValuesDigest))
    app_id TEXT NOT NULL, -- vk
    nonce BIGINT NOT NULL,
    public_values_digest TEXT NOT NULL, -- hex of bytes32
    input_data TEXT NOT NULL DEFAULT '', -- hex of bytes[], can be empty if input_url is provided
    input_url TEXT NOT NULL DEFAULT '',
    max_fee TEXT NOT NULL,
    min_stake TEXT NOT NULL,
    deadline BIGINT NOT NULL,
    created_at BIGINT NOT NULL, -- unix in seconds, blk timestamp
    processed BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (req_id)
);

CREATE INDEX IF NOT EXISTS idx_processed ON proof_request (processed);

CREATE TABLE IF NOT EXISTS my_bid (
    req_id TEXT NOT NULL,
    my_fee TEXT NOT NULL, -- can be overrided by changing bid_nonce
    bid_nonce TEXT NOT NULL,
    should_reveal_after BIGINT NOT NULL, -- unix in seconds
    should_reveal_before BIGINT NOT NULL, -- unix in seconds
    revealed BOOLEAN NOT NULL DEFAULT FALSE,
    bid_result TEXT NOT NULL DEFAULT '', -- '', 'success', 'fail' 
    proof_task_id TEXT NOT NULL DEFAULT '',
    proof_state TEXT NOT NULL DEFAULT '', -- init, generated, submitted
    proof TEXT NOT NULL DEFAULT '',
    proof_submit_tx TEXT NOT NULL DEFAULT '',
    PRIMARY KEY (req_id)
);

CREATE INDEX IF NOT EXISTS idx_revealed ON my_bid (revealed);
CREATE INDEX IF NOT EXISTS idx_bidresult_prooftaskid ON my_bid (bid_result, proof_task_id);
CREATE INDEX IF NOT EXISTS idx_proofstate ON my_bid (proof_state);

CREATE TABLE IF NOT EXISTS monitor_block (
    event TEXT NOT NULL,
    block_num BIGINT NOT NULL,
    block_idx BIGINT NOT NULL,
    restart BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (event)
);

-- NOTE: For CockroachSQL, sqlc has issue with this, need to run manually
-- GRANT ALL ON bidder.* TO bidder;