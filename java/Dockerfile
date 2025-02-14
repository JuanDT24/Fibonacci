FROM openjdk:17
WORKDIR /app  
RUN mkdir -p javat
COPY fib.java javat/fib.java   
WORKDIR /app/javat  
RUN javac fib.java  
CMD ["java", "javat.fib"]
