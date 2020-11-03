# Blackstone Microservice

Blackstone is a spaCy model and library for processing long-form, unstructured legal text. 
Here, we wrap Blackstone with a performant API layer written in Go. Communication between 
Blackstone and the API layer happens via gRPC. All of the above has been containerized to
facilitate deploying "Blackstone as a Microservice".    


## Get Started

You must have Docker installed on your machine and access to the internet. Assuming the above,
simply run: 

`docker-compose up -d`

Running the above command starts the underlying Blackstone service as well as the API layer, 
which, by default, is accessible at http://localhost:8080. Note, however, as explained below, 
the API layer currently does not support any GET requests. So, navigating to the address above
will return a `404-Not Found` error. No worries--check out Postman (https://www.postman.com/)
for easy programmatic access. 

## Routes

All of the below routes accept a POST request with a JSON body that includes a "text" property. E.g., 

```json
{
  "text": "This is the text you want to process..."
}
```

### /entities

The NER component of the Blackstone model has been trained to detect the following entity types:

| Ent        | Name           | Examples  |
| ------------- |-------------| -----:|
| CASENAME    | Case names | e.g. *Smith v Jones*, *In re Jones*, In *Jones'* case |
| CITATION      | Citations (unique identifiers for reported and unreported cases)     |   e.g. (2002) 2 Cr App R 123 |
| INSTRUMENT | Written legal instruments     |    e.g. Theft Act 1968, European Convention on Human Rights, CPR |
| PROVISION | Unit within a written legal instrument   |    e.g. section 1, art 2(3) |
| COURT | Court or tribunal   |    e.g. Court of Appeal, Upper Tribunal |
| JUDGE | References to judges |    e.g. Eady J, Lord Bingham of Cornhill |

The API layer will return a JSON response with the following shape: 

```json

{
  "data": [{"text":  "Some identified text", "label":  "CASENAME", "labelNumber":  1562316511}, ...]
}

```

### /categories 

Documentation forthcoming...

### /abbreviations 


### /compound-references 

