import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { BodkinsListComponent } from '../bodkins/bodkins-list.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, BodkinsListComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent {
  title = 'webkins';
}
