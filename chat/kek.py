from multiprocessing.pool import ThreadPool
import pandas as pd
from datetime import datetime
import tarantool

conn = tarantool.Connection('localhost', 3301)

data = pd.read_csv('chats.csv')

cnt = 1_000_000

s = 0

with ThreadPool() as p:
    for idx, row in data.iterrows():
#         if idx < 640_000:
#             continue

        try:
            p.apply(conn.call, args=(
            'insert_chat_message', row['id'], row['sender_id'], row['receiver_id'], row['text'],
            datetime.utcnow().timestamp()))
        except Exception as err:
            print(err)
            
        s += 1
        
        if s == cnt:
            break

        print(s)
