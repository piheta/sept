// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {models} from '../models';

export function Exit():Promise<void>;

export function GetAuthedUser():Promise<models.User>;

export function GetChatMessages(arg1:string):Promise<Array<models.Message>>;

export function GetChats():Promise<Array<models.Chat>>;

export function GetUser(arg1:string):Promise<models.User>;

export function GetUsers():Promise<Array<models.User>>;

export function LogOut():Promise<void>;

export function Login(arg1:string,arg2:string):Promise<models.User>;

export function Register(arg1:string,arg2:string,arg3:string):Promise<any>;

export function Search(arg1:string):Promise<Array<string>>;

export function SearchDht(arg1:string):Promise<models.User>;

export function SendMessage(arg1:string,arg2:string):Promise<Array<models.Message>>;

export function SendOffer(arg1:string):Promise<void>;
