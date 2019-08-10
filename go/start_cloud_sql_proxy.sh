#!/bin/sh

cloud_sql_proxy -instances="katana:asia-northeast2:katana-db-prod"=tcp:3306
