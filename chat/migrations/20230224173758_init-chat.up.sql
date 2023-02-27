CREATE TABLE IF NOT EXISTS chats
(
    id         UUID NOT NULL,
    sender_id UUID                     NOT NULL,
    receiver_id UUID                     NOT NULL,
    text text NOT NULL,
    created    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    PRIMARY KEY (id, sender_id)
);

CREATE INDEX IF NOT EXISTS chats_sender_receiver_created_idx ON  chats (sender_id, receiver_id, created DESC);
