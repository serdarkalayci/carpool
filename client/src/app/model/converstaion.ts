import {IContact} from "./contact";
import {IMessage} from "./message";

export interface IConversation {
  conversationid: string;
  requestername: string;
  requesterapproved: boolean;
  supplierapproved: boolean;
  requestedcapacity: number;
  requestercontact: IContact;
  suppliercontact: IContact;
  messages: IMessage[]
}

