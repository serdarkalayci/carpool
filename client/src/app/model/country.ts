import {ICity} from "./city";

export interface ICountry {
  id: string;
  name: string;
  cities: ICity[];
  ballotCities: ICity[];
}
