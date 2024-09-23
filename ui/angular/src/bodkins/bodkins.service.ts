import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Bodkin } from "./bodkin";
import { Observable, of } from "rxjs";
import {catchError, tap, map} from 'rxjs/operators'

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
  // TODO: figure out the right port number
  baseUrl = 'http://localhost:3000/bodkins'

  postBodkin(bodkin: Bodkin) {
    let url = this.baseUrl
    return this.http.post(url, bodkin, this.httpOptions).subscribe({
      next: (added) => {console.info("added: ${added}"); window.alert("Added (${added.id})!")},
      error: (e) => console.error(e),
      complete: () => console.info("create complete"),
    });
  }

  listBodkins(): Bodkin[] {
    let url = this.baseUrl
    return this.http.get(url, this.httpOptions).subscribe({
      next: (v) => {console.info("${v}")},
      error: (e) => console.error(e),
      complete: () => console.info("list complete"),
    })
  }
}
