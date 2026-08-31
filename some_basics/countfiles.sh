#!/bin/sh

# Count regular files (-type f) OR directories (-type d)
find . \( -type f -o -type d \) | wc -l
