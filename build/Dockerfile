FROM registry.access.redhat.com/ubi7/ubi-minimal:7.6

ENV OPERATOR=/usr/local/bin/deploymentConfig-operator \
    USER_UID=1001 \
    USER_NAME=deploymentConfig-operator

# install operator binary
COPY build/_output/bin/deploymentConfig-operator ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
