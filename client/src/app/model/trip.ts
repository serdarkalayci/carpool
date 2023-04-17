export interface ITrip {
  id?:string;
  countryid: string;
  origin: string;
  destination: string;
  tripdate: string;
  availableseats: number;
  stops: string[];
  note:string;
}
