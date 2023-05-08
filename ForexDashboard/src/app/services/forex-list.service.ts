import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ForexListService {

  constructor(private http: HttpClient) { }


  getForexList() {
    // TODO create type
    // return this.http.get<any>('http://localhost:8090/api/forexList');
    return this.http.get<any>('http://localhost:8092/api/currency/getAllCurrencyPair');
  }
}
