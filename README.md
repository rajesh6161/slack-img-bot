# slack-img-bot in GoLang 🤖
A Slack Bot to query google image search result and give image url of the specified query.

# Steps to be followed for setting-up Slack
1. Create your own app on Slack: https://api.slack.com/apps
2. Enable Socket Mode
![image](https://user-images.githubusercontent.com/40054161/189518761-80f565df-fe94-425e-ae73-d8a374576142.png)
3. Generate App Level Token with these permissions and this will be your SLACK_APP_TOKEN
![image](https://user-images.githubusercontent.com/40054161/189518620-114efb8e-03af-456f-a8ba-ca3d3b225bbe.png)
4. Find OAuth & Permissions sections and Generate OAuth token for you workspace, this will be your SLACK_BOT_TOKEN
5. Enable Event Subscriptions and add these permissions:
![image](https://user-images.githubusercontent.com/40054161/189518739-1a7b407c-2f39-418a-853a-b735443e9864.png)

Clone this repo add .env and paste in your secret keys and run "go mod tidy" -> "go run ."
