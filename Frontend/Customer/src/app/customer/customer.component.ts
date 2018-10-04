import { Component, OnInit } from '@angular/core';
import { Http } from '@angular/http';
import { Router } from '@angular/router';
import { Location } from '@angular/common';

@Component({
  selector: 'app-customer',
  templateUrl: './customer.component.html',
  styleUrls: ['./customer.component.css']
})
export class CustomerComponent implements OnInit {

  public customers: any;

  constructor(private http: Http, private router: Router, private location: Location) {
    this.customers = [];
  }

  ngOnInit() {
    this.location.subscribe(() => {
      this.refresh();
    });
    this.refresh();
  }

  private refresh() {
    this.http.get('http://localhost:3000/customer/')
    .subscribe(result => {
        this.customers = result.json();
    });
  }

public create() {
    this.router.navigate(['create']);
}

}
