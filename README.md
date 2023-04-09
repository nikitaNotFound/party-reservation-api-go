# Party Reservation API (Golang)

This is a Golang-based API for managing party reservations. The API is designed to be used by users who want to reserve a party or event, as well as organizers who want to manage their own parties.

## Functional Requirements

- Register using phone number
- Phone confirmation
- Registration to the parties
- Email notifications for new parties
- Email notifications for party information changes and messages organizers
- Recommendations based on user participation history
- Comments and ratings for parties and their organizers

## Non-functional Requirements

- Golang as a language
- PostgreSQL as a database
- Two microservices (one for users, another one for organizers)
- Kafka as a Message Broker for services communication
- K8s as a container orchestrator
- Azure as a cloud platform
- User app using React
- Organizers app using React


## Develop roadmap

- [ ] Database Design: 
<br />Design the database schema and create the necessary tables to store user, party, and organizer information. Use PostgreSQL as the database.
- [x] User and Organizer Microservices:
<br />Create two separate microservices - one for users and one for organizers. Each microservice should have its own API endpoint and should be responsible for handling user and organizer-related requests.
- [ ] Phone Verification:
<br />Implement phone verification during registration and party reservation using a third-party phone verification service.
- [ ] Registration to Parties:
<br />Add the ability for users to register for parties and for organizers to manage their own parties through the API.
- [ ] Email Notifications:
<br />Set up email notifications to be sent to users and organizers for new party reservations, party information changes, and other relevant updates.
- [ ] Recommendation Engine:
<br />Develop a recommendation engine that uses user participation history to suggest parties that users may be interested in attending.
- [ ] Comments and Ratings:
<br />Add a comments and ratings system to parties and their organizers, allowing users to provide feedback and ratings on their experiences.
- [ ] API Gateway:
<br />Implement an API gateway that sits in front of the user and organizer microservices and routes requests to the appropriate service.
- [ ] Kafka Integration:
<br />Integrate Kafka as the message broker for inter-service communication between the user and organizer microservices.
- [ ] User and Organizer Apps:
<br />Develop separate React-based apps for users and organizers to interact with the API and manage their reservations/parties.
- [ ] Containerization with Kubernetes:
<br />Containerize the microservices using Docker and deploy them on Kubernetes for scalability and availability.
- [ ] Cloud Platform Deployment:
<br />Deploy the application to Azure, leveraging its capabilities for cloud-based infrastructure.
