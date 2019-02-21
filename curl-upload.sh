#!/bin/bash

curl \
-F "image=@dogo.jpg" \
-F "image=@dogo_clone.jpg" \
-F "image=@dogo2.jpg" \
-H "Content-Type: multipart/form-data" \
localhost:8081/data