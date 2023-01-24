import hashlib
import random
import uuid

import pandas as pd
from faker import Faker

FAKE = Faker(locale='ru_RU')
COUNT = 1_000_000
DATA = {
    'id': [],
    'last_name': [],
    'first_name': [],
    'age': [],
    'gender': [],
    'biography': [],
    'city': [],
    'password': []
}
COUNT_HT_PARAMS = 10_000
HT_PARAMS = {
    'last_name': [],
    'first_name': [],
}


def generate_one():
    id = uuid.uuid4()
    last_name = FAKE.last_name()
    first_name = FAKE.first_name()
    age = random.randint(10, 99)
    gender = random.choice(['male', 'female', 'other'])
    biography = FAKE.text()
    city = FAKE.city_name()
    password = hashlib.sha256(str(random.randint(1, 1_000_000)).encode()).hexdigest()

    DATA['id'].append(id)
    DATA['last_name'].append(last_name)
    DATA['first_name'].append(first_name)
    DATA['age'].append(age)
    DATA['gender'].append(gender)
    DATA['biography'].append(biography)
    DATA['city'].append(city)
    DATA['password'].append(password)


def main():
    print("=== START GENERATE DATA ===")
    for _ in range(COUNT):
        generate_one()

    df = pd.DataFrame(DATA)
    df.to_csv('people.csv', index=False)

    print("=== START GENERATE HT PARAMS ===")
    df = df.iloc[:COUNT_HT_PARAMS]
    for tpl in df.to_records(index=False):
        HT_PARAMS['last_name'].append(tpl[1][:random.randint(1, 4)])
        HT_PARAMS['first_name'].append(tpl[2][:random.randint(1, 4)])

    random.shuffle(HT_PARAMS['last_name'])
    random.shuffle(HT_PARAMS['first_name'])

    df = pd.DataFrame(HT_PARAMS)
    df.to_csv('people_params.csv', index=False, header=False)


if __name__ == '__main__':
    main()
