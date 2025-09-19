#!/usr/bin/env bash

pushd "cmd/"

fyne package -os android \
  --app-id com.hazardous_sun.transaction_tracker \
  --name "Transaction Tracker" \
  --icon ../Icon.png \
  --release

popd
