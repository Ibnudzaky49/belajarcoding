#!/bin/bash

host=$(zenity --title="insert hostname" --entry="password")
pass=$(zenity --title="insert password" --entry="Name") &&
  echo "$pass" | sudo -S sh -c "echo $host > /etc/hostname"
