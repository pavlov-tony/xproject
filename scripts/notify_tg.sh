#! /bin/bash

NEWLINE=$'\n'
curl -X POST https://api.telegram.org/bot657121652:AAEz_Uy77hcSBx5Brd_DA3F62J6yKmZBCmU/sendMessage -d chat_id=-1001361347769 \
-d "text=branch:${TRAVIS_BRANCH}${NEWLINE}commit:${TRAVIS_COMMIT}${NEWLINE}${TRAVIS_COMMIT_MESSAGE}${NEWLINE}tests return code: ${TRAVIS_TEST_RESULT}"