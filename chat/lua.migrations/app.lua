#!/usr/bin/env tarantool
box.cfg {}

-- add 512Mb
box.cfg({memtx_memory = box.cfg.memtx_memory + 512 * 2^20})

box.once("schema", function()
    -- Define the "chats" space
    box.schema.create_space('chats', {
        format = {
            {name = 'id',         type = 'string'},
            {name = 'sender_id',  type = 'string'},
            {name = 'receiver_id',type = 'string'},
            {name = 'text',       type = 'string'},
            {name = 'created',    type = 'scalar'},
        },
        if_not_exists = true,
    })

    -- Define the primary key
    box.space.chats:create_index('primary', {
        parts = {'id'},
        if_not_exists = true,
    })

    -- Define the secondary index
    box.space.chats:create_index('sender_receiver', {
        parts = {'sender_id', 'receiver_id', 'created'},
        if_not_exists = true,
    })

end)

function select_chats(sender_id, receiver_id)
    local chats = box.space.chats.index.sender_receiver:select({sender_id, receiver_id}, {iterator = 'EQ'})
    table.sort(chats, function(a, b) return a[5] > b[5] end)
    return chats
end

function insert_chat_message(id, sender_id, receiver_id, text, created)
   box.space.chats:insert{ id, sender_id, receiver_id, text, created }
end

function hello()
   return "hello"
end