import uuid
import jwt

import pandas as pd
from faker import Faker
import datetime

FAKE = Faker(locale='ru_RU')
COUNT_FRIENDS_PAIR = 10_000
COUNT_CHAT_SIZE = 20
DATA = {
    'id': [],
    'sender_id': [],
    'receiver_id': [],
    'text': []
}
COUNT_HT_PARAMS = 20_000
HT_PARAMS = {
    'token': [],
    'receiver_id': [],
}


def generate_one(sender_id, receiver_id):
    id = uuid.uuid4()
    text = FAKE.text()

    DATA['id'].append(id)
    DATA['sender_id'].append(sender_id)
    DATA['receiver_id'].append(receiver_id)
    DATA['text'].append(text)


def create_token(sender_id):
    return jwt.encode({
        "iss": "backend.auth.service",
        "exp": datetime.datetime.now() + datetime.timedelta(days=1),
        "user_id": sender_id,
        "token_type": "access"
    }, "kek", algorithm="HS256")


def main():
    print("=== START GENERATE DATA ===")
    for _ in range(COUNT_FRIENDS_PAIR):
        sender_id = uuid.uuid4()
        for _ in range(COUNT_CHAT_SIZE):
            receiver_id = uuid.uuid4()
            for _ in range(COUNT_CHAT_SIZE):
                generate_one(sender_id, receiver_id)

    df = pd.DataFrame(DATA)
    df.to_csv('chats.csv', index=False)

    print("=== START GENERATE HT PARAMS ===")
    df = df.iloc[:COUNT_HT_PARAMS]
    for tpl in df.to_records(index=False):
        HT_PARAMS['token'].append(create_token(str(tpl[1])))
        HT_PARAMS['receiver_id'].append(tpl[2])

    df = pd.DataFrame(HT_PARAMS)
    df.to_csv('chat_params.csv', index=False, header=False)


if __name__ == '__main__':
    main()
