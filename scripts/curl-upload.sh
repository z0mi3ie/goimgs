#!/bin/bash

curl \
-F "image=@dogo.jpg" \
-H "Content-Type: multipart/form-data" \
localhost:8081/images
