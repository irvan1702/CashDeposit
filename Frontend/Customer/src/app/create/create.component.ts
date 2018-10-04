import { Component, OnInit } from '@angular/core';
import { Http } from '@angular/http';
import { Location } from '@angular/common';

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.css']
})
export class CreateComponent implements OnInit {
  public customer: any;

  public constructor(private location: Location, private http: Http) {
    this.customer = {
        'customer_name': '',
        'account_number': '',
        'cash_deposit': 0
    };
}

public ngOnInit() { }

public save() {
    if (this.customer.customer_name && this.customer.account_number && this.customer.cash_deposit) {
        this.http.post('http://localhost:3000/customer/update', JSON.stringify(this.customer))
            .subscribe(result => {
                console.log(this.customer);
                console.log(result);
                this.location.back();
            });
    }
}

}
