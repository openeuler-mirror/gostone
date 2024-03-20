import os
import sys

from celery import Celery
from kombu import Queue

app = Celery('tasks', broker='amqp://jenkins:Jenkins@123QWE@172.28.8.12//')

app.conf.task_queues = (
    Queue('jenkins', routing_key='jenkins'),
    Queue('steel', routing_key='steel'),
)
app.conf.task_acks_late = True


@app.task(name='jenkins_rpm_sync.RPM')
def RPM(url, sub_path='CT-dev/gostack'):
    local_path = "/data/repo/yum/ctyun/centos/7/cloud/x86_64"
    cmd = "cd %s/%s && curl -O %s" % (local_path, sub_path, url)
    os.system(cmd)
    update_repo_jenkins = "createrepo %s/%s" % (local_path, sub_path)
    os.system(update_repo_jenkins)


RPM_package_name_url = sys.argv[1]
yum_repo_sub_path = sys.argv[2]
RPM.apply_async([RPM_package_name_url, yum_repo_sub_path], queue='jenkins')
