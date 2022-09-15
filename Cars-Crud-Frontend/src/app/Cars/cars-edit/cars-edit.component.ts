import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params, Router } from '@angular/router';
import { FormBuilder, Validators } from '@angular/forms';
import {CarsService} from '../../services/cars.service'

@Component({
  selector: 'app-cars-edit',
  templateUrl: './cars-edit.component.html',
  styleUrls: ['./cars-edit.component.scss']
})
export class CarsEditComponent implements OnInit {

  id: any;
  car: any;
  model: any;
  year: any;

  checkoutForm = this.formBuilder.group({
    Id: ['', Validators.required],
    Car: ['', Validators.required],
    Model: ['', Validators.required],
    Year: ['', Validators.required],

  });

  constructor(private route: ActivatedRoute,
              private formBuilder: FormBuilder,
              private carService: CarsService,
              private router: Router) { }

  ngOnInit(): void {
    this.id = this.route.snapshot.paramMap.get('id')
    this.car = this.route.snapshot.paramMap.get('name')
    this.model = this.route.snapshot.paramMap.get('model')
    this.year = this.route.snapshot.paramMap.get('Year')



    this.checkoutForm.controls['Id'].patchValue(this.id);
    this.checkoutForm.controls['Car'].patchValue(this.car);
    this.checkoutForm.controls['Model'].patchValue(this.model);
    this.checkoutForm.controls['Year'].patchValue(this.year);
  }

  onSubmit(){

    let Car = this.checkoutForm.value.Car;
    let Model = this.checkoutForm.value.Model;
    let Year = this.checkoutForm.value.Year;
    let id = this.checkoutForm.value.Id

    let car = { id ,Car, Model, Year};

    console.log(car);

    if(Car === '' || Model === '' || Year ===  '' ){
      alert('Please fill all inputs')
      return;
    }
    //call service to insert valus of database
    this.carService.editCar(car).subscribe(result => {
      alert(result)
      this.checkoutForm.reset();
      this.router.navigate(['/cars']);

    })

  }

}
