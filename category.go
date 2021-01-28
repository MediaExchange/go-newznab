/*
   Copyright 2021 MediaExchange.io

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package newznab

type Category struct {
	id string
}

var (
	Console            = Category{id: "1000"}
	Console_NDS        = Category{id: "1010"}
	Console_PSP        = Category{id: "1020"}
	Console_Wii        = Category{id: "1030"}
	Console_Xbox       = Category{id: "1040"}
	Console_Xbox360    = Category{id: "1050"}
	Console_WiiWare    = Category{id: "1060"}
	Console_Xbox360DLC = Category{id: "1070"}

	Movies         = Category{id: "2000"}
	Movies_Foreign = Category{id: "2010"}
	Movies_Other   = Category{id: "2020"}
	Movies_SD      = Category{id: "2030"}
	Movies_HD      = Category{id: "2040"}
	Movies_UHD     = Category{id: "2045"}
	Movies_BluRay  = Category{id: "2050"}
	Movies_3D      = Category{id: "2060"}

	Audio           = Category{id: "3000"}
	Audio_MP3       = Category{id: "3010"}
	Audio_Video     = Category{id: "3020"}
	Audio_Audiobook = Category{id: "3030"}
	Audio_Lossless  = Category{id: "3040"}

	PC                = Category{id: "4000"}
	PC_0Day           = Category{id: "4010"}
	PC_ISO            = Category{id: "4020"}
	PC_Mac            = Category{id: "4030"}
	PC_Mobile_Other   = Category{id: "4040"}
	PC_Games          = Category{id: "4050"}
	PC_Mobile_IOS     = Category{id: "4060"}
	PC_Mobile_Android = Category{id: "4070"}

	TV             = Category{id: "5000"}
	TV_Foreign     = Category{id: "5020"}
	TV_SD          = Category{id: "5030"}
	TV_HD          = Category{id: "5040"}
	TV_UHD         = Category{id: "5045"}
	TV_Other       = Category{id: "5050"}
	TV_Sport       = Category{id: "5060"}
	TV_Anime       = Category{id: "5070"}
	TV_Documentary = Category{id: "5080"}

	XXX        = Category{id: "6000"}
	XXX_DVD    = Category{id: "6010"}
	XXX_WMV    = Category{id: "6020"}
	XXX_XviD   = Category{id: "6030"}
	XXX_x264   = Category{id: "6040"}
	XXX_Pack   = Category{id: "6050"}
	XXX_ImgSet = Category{id: "6060"}
	XXX_Other  = Category{id: "6070"}

	Books        = Category{id: "7000"}
	Books_Mags   = Category{id: "7010"}
	Books_EBook  = Category{id: "7020"}
	Books_Comics = Category{id: "7030"}

	Other      = Category{id: "8000"}
	Other_Misc = Category{id: "8010"}
)
