# Functions

## Get

Chat dialogue API is an API that provides a wide range of conversations and "shiritori" for natural speech of a user.

#### Params

- **requestBody** - **RequestBody** - RequestBody structure

#### API Command

**Get**

- **utt** - string - Utterance of the user (255 characters or less)  ex. こんにちは
- **context** - string -  Context id  ex. aaabbbccc111222333
- **nickname** - string - Nickname of the user (10 characters or less)  ex. 光
- **nicknameY** - string - Kana of the nickname of the user (20 characters or less)  ex. ヒカリ
- **sex** - string - Sex of the user (男 or 女)  ex. 女
- **bloodtype** - string - Blood type of user (A、B、AB、O)  ex. A
- **birthdateY** - string - Year of the user's birthday (1-)  ex. 1997
- **birthdateM** - string - Month of the user's birthday (1-12)  ex. 5
- **birthdateD** - string - Day of the user's birthday (1-31)  ex. 30
- **age** - string - Age of the user (1-999)  ex. 16
- **constellations** - string - Constellations of user (牡羊座,牡牛座,双子座,蟹座,獅子座,乙女座,天秤座,蠍座,射手座,山羊座,水瓶座,魚座)  ex. 双子座
- **place** - string - Regional information of the user (https://devsite-pro.s3.amazonaws.com/contents_file/Dialogue_API_spec_v1.0.2.pdf)  ex. 東京
- **mode** - string - Mode of dialogue (dialog,srtr)  ex. dialog
