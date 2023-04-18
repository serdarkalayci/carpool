import {Component, Input} from '@angular/core';
import {ITrip} from "../model/trip";

@Component({
  selector: 'cp-tripdetails',
  templateUrl: './tripdetails.component.html',
  styleUrls: ['./tripdetails.component.css']
})
export class TripdetailsComponent {
  @Input() currentTrip: ITrip |undefined;

}
