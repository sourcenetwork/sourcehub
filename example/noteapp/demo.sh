#!/usr/bin/bash

set -e

NOTEBIN="./build/noteapp"
FAUCET="./scripts/faucet.sh"


read -p "Generating Accounts for Alice, Bob and Charlie. Press Enter to continue"

$NOTEBIN gen-account > ./tmp
ALICE=$(head -n 1 ./tmp)
ALICE_PRIV="alice.priv"
tail -n 1 ./tmp > $ALICE_PRIV
$FAUCET $ALICE

$NOTEBIN gen-account > ./tmp
BOB=$(head -n 1 ./tmp)
BOB_PRIV="bob.priv"
tail -n 1 ./tmp > $BOB_PRIV
sleep 4
$FAUCET $BOB

$NOTEBIN gen-account > ./tmp
CHARLIE=$(head -n 1 ./tmp)
CHARLIE_PRIV="charlie.priv"
tail -n 1 ./tmp > $CHARLIE_PRIV
sleep 4
$FAUCET $CHARLIE

rm ./tmp


echo "Generated Accounts"
echo "Alice: $ALICE"
echo "Bob: $BOB"
echo "Chalie: $CHARLIE"


read -p "2. Create Alice's Notes. Press Enter to continue"
PRIV_ID=$($NOTEBIN new-note $ALICE_PRIV 'alice.priv.txt' "Alice's secret")
sleep 4
PUB_ID=$($NOTEBIN new-note $ALICE_PRIV 'alice.txt' 'hello from Alice')
echo "Created Notes alice.txt ($PUB_ID) and alice.priv.txt ($PRIV_ID)"


read -p "3. List Alice's local Notes. Press Enter to continue" </dev/tty
$NOTEBIN list-local-notes $ALICE_PRIV 2> /dev/null


read -p "4: List Bob's Notes (should not be able to see any document). Press Enter to continue."
$NOTEBIN list-notes $BOB_PRIV 2> /dev/null


read -p "5: Make Alice share alice.txt ($PUB_ID) with Bob. Press Enter to continue."
$NOTEBIN share $ALICE_PRIV $PUB_ID $BOB 2> /dev/null


read -p "6: List Bob's Notes (should see alice.txt). Press Enter to continue."
$NOTEBIN list-notes $BOB_PRIV


read -p "7: List Charlie's Notes (should list nothing). Press Enter to continue."
$NOTEBIN list-notes $CHARLIE_PRIV 2> /dev/null
