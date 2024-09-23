import { Component, OnInit } from '@angular/core';
import { BodkinsService } from './bodkins.service'
import {Router} from '@angular/router'
import {Bodkin} from "./bodkin"

@Component({
  selector: 'app-bodkins-create',
  standalone: true,
  imports: [],
  templateUrl: './bodkins-create.component.html',
  styleUrl: './bodkins-create.component.css'
})
export class BodkinsCreateComponent implements OnInit {

  name = ''

  constructor(private service: BodkinsService, private router: Router) { }

  ngOnInit(): void {
  }

  // doesn't seem right.. but maybe
  createBodkin() {
    // id is required but for create it will get ignored.
    let bodkin = {id: 0, name: this.name}

    this.service.postBodkin(bodkin);
  }

}
