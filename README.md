# Make Java Keystore (from PEMs)

[![Build Status](https://travis-ci.org/bsycorp/make-jks.svg?branch=master)](https://travis-ci.org/bsycorp/make-jks)

A statically linked keytool binary that is extended to take PEM-encoded certificates (like /etc/ssl/certs) and create a Java Keystore out of it (JKS). This is useful for java processes running in docker containers where `/etc/ssl/certs` is mounted into the container, but the Java keystore isn't updated to match. All the original Java keytool functionality also still works, listing etc.

## Run

`make-jks -import -dir /etc/ssl/certs -keystore $JAVA_HOME/lib/security/cacerts -storepass changeit`

`make-jks -list -keystore $JAVA_HOME/lib/security/cacerts -storepass changeit`