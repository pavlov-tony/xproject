#! /bin/bash

curl -X POST https://api.telegram.org/bot657121652:AAEz_Uy77hcSBx5Brd_DA3F62J6yKmZBCmU/sendMessage -d chat_id=-1001361347769 \
-d "text=${TRAVIS_BRANCH} (${TRAVIS_COMMIT_MESSAGE}) tests return code: ${TRAVIS_TEST_RESULT}"