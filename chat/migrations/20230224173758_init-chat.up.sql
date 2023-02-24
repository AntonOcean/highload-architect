CREATE TABLE IF NOT EXISTS chats
(
    id         UUID PRIMARY KEY,
    sender_id UUID                     NOT NULL,
    receiver_id UUID                     NOT NULL,
    text text NOT NULL,
    created    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);