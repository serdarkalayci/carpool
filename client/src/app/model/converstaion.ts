import {IContact} from "./contact";

export interface IConversation {
  conversationid: string;
  requestername: string;
  requesterapproved: boolean;
  supplierapproved: boolean;
  requestedcapacity: 3;
  requestercontact: IContact;
  suppliercontact: IContact;
  messages: string[]
}
