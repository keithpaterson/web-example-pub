import { Component, OnInit } from '@angular/core';
import { BodkinsService } from './bodkins.service'
import { Router } from '@angular/router'
import { Bodkin } from "./bodkin"
import { NgFor } from '@angular/common';

@Component({
  selector: 'app-bodkins-list',
  standalone: true,
  imports: [NgFor],
  templateUrl: './bodkins-list.component.html',
  styleUrl: './bodkins-list.component.css'
})
export class BodkinsListComponent {
  bodkins: Bodkin[] = [];
  
  constructor(private service: BodkinsService, private router: Router) {
    this.listBodkins()
  }

  ngOnInit(): void {
    this.listBodkins()
  }

  listBodkins() {
    this.bodkins = this.service.listBodkins()
  }
}
