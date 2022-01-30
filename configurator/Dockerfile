FROM python:3

RUN apt-get update -y && apt-get install -y npm
# Upgrade to the latest version of Node.js
RUN npm install n -g && n latest && hash -r && npm install -g npm@8.3.2 

RUN mkdir /app
COPY ./app/package.json ./app/requirements.txt /app 

RUN pip install --upgrade pip && pip install -r /app/requirements.txt

RUN apt-get install -y vim
WORKDIR /app
RUN npm install
# This will complain about 8 packages
COPY ./root_bash_history /root/.bash_history


COPY ./app /app 

#CMD [ "npm", "start" ]
CMD [ "sleep", "5000" ]
