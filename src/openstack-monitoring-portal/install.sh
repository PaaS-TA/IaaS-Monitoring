#!/usr/bin/env bash

. .envrc
go get github.com/tedsuo/ifrit

#rata
go get github.com/tedsuo/rata

#InfluxDB Library Download
go get github.com/influxdata/influxdb/client/v2

#go client for openstack
go get github.com/rackspace/gophercloud

#Mysql Driver and Orm Library Download
go get github.com/go-sql-driver/mysql
go get github.com/jinzhu/gorm
go get github.com/cihub/seelog
go get github.com/monasca/golang-monascaclient/monascaclient
go get github.com/gophercloud/gophercloud/
go get github.com/alexedwards/scs

# elastic search
go get gopkg.in/olivere/elastic.v3

#monasca client Bug Fix Src
cp ./lib-bugFix-src/alarm_definitions.go ./src/github.com/monasca/golang-monascaclient/monascaclient
cp ./lib-bugFix-src/notifications.go ./src/github.com/monasca/golang-monascaclient/monascaclient
cp ./lib-bugFix-src/alarms.go ./src/github.com/monasca/golang-monascaclient/monascaclient