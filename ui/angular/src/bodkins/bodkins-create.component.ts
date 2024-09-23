import { Component, OnInit } from '@angular/core';
import { BodkinsService } from './bodkins.service'
import {Router} from '@angular/router'
import {Bodkin} from "./bodkin"

@Component({
  selector: 'app-bodkins',
  standalone: true,
  imports: [],
  templateUrl: './bodkins.component.html',
  styleUrl: './bodkins.component.css'
})
export class BodkinsComponent implements OnInit {

  name = ''

  constructor(private service: BodkinsService, private router: Router) { }

  ngOnInit(): void {
  }

  createBodkin() {
    // id is required but for create it will get ignored.
    let bodkin = {id: 0, name: this.name}

    this.service.postBodkin(bodkin);
  }

}
