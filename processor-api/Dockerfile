# stage 1: build 
FROM python:3-alpine AS builder

WORKDIR /processor-api 

COPY requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

COPY . .

VOLUME [ "/processor-api/summaries" ]

# hugging face token should be given while running the container
# -e HUGGING_FACE_TOKEN="token"  
# ENV HUGGING_FACE_TOKEN=""

EXPOSE 5000

CMD [ "python3", "main.py"]

# stage 1: build 
# FROM python:3-alpine AS builder

# WORKDIR /processor-api 

# RUN python3 -m venv /opt/venv
# ENV PATH="/opt/venv/bin:$PATH"

# COPY requirements.txt .

# RUN pip install --no-cache-dir -r requirements.txt


# # stage 2: final image
# FROM python:3-alpine

# WORKDIR /processor-api

# EXPOSE 5000

# # copy only installed dependencies
# COPY --from=builder /opt/venv /opt/venv
# ENV PATH="/opt/venv/bin:$PATH"

# COPY . .

# VOLUME [ "/processor-api/summaries" ]

# # hugging face token should be given while running the container
# # -e HUGGING_FACE_TOKEN="token"  
# # ENV HUGGING_FACE_TOKEN=""

# CMD [ "python3", "main.py"]