FROM python:3.6-alpine
COPY ./code /code
WORKDIR /code
RUN pip install redis -i https://pypi.douban.com/simple \
    && pip install flask -i https://pypi.douban.com/simple
CMD ["python", "app.py"]

