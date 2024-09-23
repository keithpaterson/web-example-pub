import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Bodkin } from "./bodkin";
import { environment } from "../environment/environment"
//import { Observable, of } from "rxjs";
//import {catchError, tap, map} from 'rxjs/operators'

@Injectable({
  providedIn: 'root'
})
export class BodkinsService {

  constructor(private http: HttpClient) { }

  httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json'
    })
  }

  baseUrl = 'http://localhost:' + environment.bodkinsPort + '/bodkins'

  postBodkin(bodkin: Bodkin) {
    console.log("post new bodkin at port " + environment.bodkinsPort)
    let url = this.baseUrl
    return this.http.post(url, bodkin, this.httpOptions).subscribe({
      next: (added) => {console.info("added: ${added}"); window.alert("Added (${added.id})!")},
      error: (e) => console.error(e),
      complete: () => console.info("create complete"),
    });
  }

  listBodkins(): Bodkin[] {
    console.log("list bodkin at port " + environment.bodkinsPort)
    let url = this.baseUrl
    let bodkins: Bodkin[] = [];
    this.http.get(url, this.httpOptions).subscribe({
      next: (v) => {console.info(v); /* TODO: capture the array */},
      error: (e) => console.error(e),
      complete: () => console.info("list complete"),
    });
    return bodkins;
  }
}
