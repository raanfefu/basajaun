import { Component } from '@angular/core';
import { NgbModal, ModalDismissReasons } from '@ng-bootstrap/ng-bootstrap';
import { Policy, Term, Statement, Func, Bundle } from './models'
import { catchError, throwError} from 'rxjs';
import Ajv, { JSONSchemaType } from "ajv"
import { HttpService } from './http.service';

import { HttpErrorResponse } from '@angular/common/http';


const ajv = new Ajv()
const schema: JSONSchemaType<Bundle> = {
  type: "object",
  properties: {
    policies: {
      type: "array",
      items: {
        type: "object",
        properties: {
          uri_resource: {
            type: "string"
          },
          action: {
            type: "array",
            items: {
              type: "string",
              enum: ["PUT",
                "GET",
                "POST",
                "DELETE",
                "HEAD",
                "CONNECT",
                "OPTIONS",
                "TRACE",
                "PATCH"]
            },
            minItems: 1
          },
          statements: {
            type: "array",
            items: {
              type: "object",
              properties: {
                type: {
                  type: "string",
                  enum: [
                    "body",
                    "header",
                    "claims",
                    "const",
                    "data",
                    "query",
                    "path",
                    "cache"
        
                    
                  ]
                },
                uri: {
                  type: "string",
                  pattern: "^[A-Za-z-_\\./]+[A-Za-z0-9-_\\./]*$"
                },
                func: {
                  type: "object",
                  properties: {
                    type: {
                      type: "string",
                      enum: [
                        "equals",
                        "not-equals",
                        "in"
                      ]
                    },
                    term: {
                      type: "object",
                      properties: {
                        type: {
                          type: "string",
                          enum: [
                            "body",
                            "header",
                            "const",
                            "data",
                            "query",
                            "path",
                            "cache",
                            "claims"

                          ]
                        },
                        uri: {
                          type: "string", 
                          pattern: "^[A-Za-z-_\\./]+[A-Za-z0-9-_\\./]*$"
                        }
                      },
                      required: ["type", "uri"]

                    }
                  },
                  required: ["type", "term"]
                }
              },
              required: [
                "type",
                "uri",
                "func"
              ]
            },
            minItems: 1
          }
        },
        required: [
          "uri_resource",
          "action",
          "statements"
        ]
      },
      minItems: 1
    }
  },
  required: [
    "policies"
  ]
}

const validate = ajv.compile(schema)

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'gui-bdle-plcy';
  disabledButton = false
  response : any = {}

  settings : any = {}

  namebundle : string = "sample-authz"
  publish = false
  refresh = false
  messagePublish = ""

  actionseleted: Array<any> = [];
  actions = [
    { name: "PUT", id: 1     },
    { name: "GET", id: 2     },
    { name: "POST", id: 3    },
    { name: "DELETE", id: 4  },
    { name: "HEAD", id: 3    },
    { name: "CONNECT", id: 3 },
    { name: "OPTIONS", id: 3 },
    { name: "TRACE", id: 3   },
    { name: "PATCH", id: 3   },
  ];

  onChange(event: any, target: Policy) {
    let value = event.target.id
    let isChecked = event.target.checked
    console.log(`${value} = ${isChecked}`)
    if (isChecked) {
      target.action.push(value);
    } else {
      let index = target.action.indexOf(value);
      target.action.splice(index, 1);
    }
  }

  public bundle: Bundle = {
    policies: [
    ]
  }

  constructor(private modalService: NgbModal, private httpService: HttpService) {
  }


 
  public addPolicy() {
    this.response = {}
    this.bundle.policies.push(
      {
        "uri_resource": "/api/v1/order",
        "action": [],
        "statements": [
          {
            "type": "body",
            "uri": "spec.user.id",
            "func": {
              "type": "equals",
              "term": {
                "type": "header",
                "uri": "x-user-id"
              }
            }
          }
        ]
      }
    );
  }

  public addStmt(index: number) {
    this.bundle.policies[index].statements.push(
      {
        "type": "body",
        "uri": "spec.user.id",
        "func": {
          "type": "equals",
          "term": {
            "type": "header",
            "uri": "x-user-id"
          }
        }
      }
    );
  }

  public delStmt(indexPolicy: number, indexStmt: number) {
    this.bundle.policies[indexPolicy].statements.splice(indexStmt, 1);
  }

  public delPolicy(indexPolicy: number) {
    this.bundle.policies.splice(indexPolicy, 1);
  }

  public evalPolicy() {
    return validate(this.bundle)
  }

  public format : boolean = false
  public toggle(){
    this.format = !this.format
  }


  public closeResult: string = "";

  

  public openModelSettings(content: any){
    this.messagePublish = ""
    this.modalService.open(content, {
      windowClass: "modal-lg",
      ariaLabelledBy: 'modal-basic-title'
    }).result.then((result) => {
      this.closeResult = `Closed with: ${result}`;
    }, (reason) => {
      this.closeResult = reason;
    });
  }

  public close(reason: number) {
    if(this.modalService.hasOpenModals()){
      this.modalService.dismissAll(reason) 
    }
  }



  private getDismissReason(reason: any): string {
    if (reason === ModalDismissReasons.ESC) {
      return 'by pressing ESC';
    } else if (reason === ModalDismissReasons.BACKDROP_CLICK) {
      return 'by clicking on a backdrop';
    } else {
      return  `with: ${reason}`;
    }
  }

  public callPublish(){
    if (validate(this.bundle)) {
      this.publish = true
      this.httpService.postPublish({
        name: this.namebundle,
        rule: this.bundle
      })
      .subscribe((data) => {
          this.publish = false 
          this.messagePublish = `uploaded ${this.namebundle}`
      });
      
   
    }
  }

  public openModelPublish(content: any){
    this.modalService.open(content, {
      windowClass: "modal",
      ariaLabelledBy: 'modal-basic-title'
    }).result.then((result) => {
      this.closeResult = `Closed with: ${result}`;
    }, (reason) => {
      this.closeResult = reason;
    });
  }
 

  public openSettings(content: any) {
   
    this.httpService.getSettings()
        .subscribe((data) => {
          this.settings = data
          this.openModelSettings(content)
        });
  }

  
  public callPolicy() {
    if (validate(this.bundle)) {
      this.refresh = true
        this.httpService.postData(this.bundle)
        .subscribe((data) => {
          this.refresh = false
            this.response = data
        });
     
    }
  }

}
