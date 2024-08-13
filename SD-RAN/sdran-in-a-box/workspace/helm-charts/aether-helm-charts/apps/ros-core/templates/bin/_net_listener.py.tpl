# Copyright 2019-present Open Networking Foundation
#
# SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

import argparse, socket
from socket import gethostbyaddr
from datetime import datetime

sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
sock.bind(('0.0.0.0', {{ .Values.config.droneNetListener.port }}))
print('Listening at {}'.format(sock.getsockname()))
while True:
    data, address = sock.recvfrom(65535)
    text = data.decode('ascii')
    print('The client at {} says {!r}'.format(address, text))
