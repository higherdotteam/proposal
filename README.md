
# proposal

This program runs a bot named "proposal" who joins your slack team. Users can direct message proposal and say "hello" or anything to start the bot going. The bot will ask the user 18 questions and record their answers. When done, a new private channel is created with the questions and answers.

# setup

Step 1.
https://my.slack.com/apps/manage/custom-integrations

From ^ that link make a new bot integration and name your bot "proposal"

Copy the bot's API token to your clipboard and add it to your .bash_profile like:

```
export SLACK_PROPOSAL_BOT="xoxb-123456789-ORxt5Xg2yqV8jvDS5fVW9AAZ"
expoort SLACK_PROPOSAL_ADMIN="xoxb-123456789-56xt5Xg2yqV8jvDS5fVW9ABY"
```

You can get your admin's token via:

https://get.slack.help/hc/en-us/articles/215770388-Create-and-regenerate-API-tokens

https://api.slack.com/docs/oauth-test-tokens

http://higher.team/tokens
