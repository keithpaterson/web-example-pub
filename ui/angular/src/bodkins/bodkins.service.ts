import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Bodkin } from "./bodkin";
import { environment } from "../environment/environment"
import { Observable } from 'rxjs';
//import { Observable, of } from "rxjs";
//import {catchError, tap, map} from 'rxjs/operators'

@Injectable({
  providedIn: 'root'
})
export class BodkinsService {

  constructor(private http: HttpClient) { }

  private httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json'
    })
  }

  private baseUrl = 'http://localhost:' + environment.bodkinsPort + '/bodkins'

  postBodkin(bodkin: Bodkin): Observable<Object> {
    console.log("post new bodkin at port " + environment.bodkinsPort)
    let url = this.baseUrl
    return this.http.post(url, bodkin, this.httpOptions);
    /*
    let added: Bodkin = {id: -1, name: ''};
    this.http.post(url, bodkin, this.httpOptions).subscribe({
      next: (v) => {added = v as Bodkin; console.info("added: ${v}"); window.alert("Added (${added.id})!");},
      error: (e) => console.error(e),
      complete: () => console.info("create complete"),
    });
    return added;
    */
  }

  listBodkins(): Observable<Object> {
    console.log("list bodkin at port " + environment.bodkinsPort)
    let url = this.baseUrl;
    return this.http.get(url, this.httpOptions);
    /*
    let bodkins: Bodkin[] = [];
    this.http.get(url, this.httpOptions).subscribe({
      next: (v) => {bodkins = v as Bodkin[]; console.info("next: " + v + " (" + bodkins + ")");},
      error: (e) => console.error(e),
      complete: () => console.info("list complete (" + bodkins + ")"),
    });
    console.info("returning: " + bodkins);
    return bodkins;
    */
  }
}
