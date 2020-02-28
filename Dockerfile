FROM alpine
ADD jasontest-srv /jasontest-srv
ENTRYPOINT [ "/jasontest-srv" ]
