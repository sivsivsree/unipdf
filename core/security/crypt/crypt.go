//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package crypt ;import (_d "crypto/aes";_e "crypto/cipher";_g "crypto/md5";_c "crypto/rand";_fg "crypto/rc4";_ff "fmt";_fb "github.com/unidoc/unipdf/v3/common";_ed "github.com/unidoc/unipdf/v3/core/security";_b "io";);func init (){_gdg ("\u0041\u0045\u0053V\u0032",_ec )};func (filterIdentity )EncryptBytes (p []byte ,okey []byte )([]byte ,error ){return p ,nil };

// NewFilter creates CryptFilter from a corresponding dictionary.
func NewFilter (d FilterDict )(Filter ,error ){_ece ,_cga :=_ca (d .CFM );if _cga !=nil {return nil ,_cga ;};_gd ,_cga :=_ece (d );if _cga !=nil {return nil ,_cga ;};return _gd ,nil ;};func (filterAES )DecryptBytes (buf []byte ,okey []byte )([]byte ,error ){_ge ,_cb :=_d .NewCipher (okey );if _cb !=nil {return nil ,_cb ;};if len (buf )< 16{_fb .Log .Debug ("\u0045R\u0052\u004f\u0052\u0020\u0041\u0045\u0053\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0062\u0075\u0066\u0020\u0025\u0073",buf );return buf ,_ff .Errorf ("\u0041\u0045\u0053\u003a B\u0075\u0066\u0020\u006c\u0065\u006e\u0020\u003c\u0020\u0031\u0036\u0020\u0028\u0025d\u0029",len (buf ));};_db :=buf [:16];buf =buf [16:];if len (buf )%16!=0{_fb .Log .Debug ("\u0020\u0069\u0076\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (_db ),_db );_fb .Log .Debug ("\u0062\u0075\u0066\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );return buf ,_ff .Errorf ("\u0041\u0045\u0053\u0020\u0062\u0075\u0066\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069p\u006c\u0065\u0020\u006f\u0066 \u0031\u0036 \u0028\u0025\u0064\u0029",len (buf ));};_fc :=_e .NewCBCDecrypter (_ge ,_db );_fb .Log .Trace ("A\u0045\u0053\u0020\u0044ec\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );_fb .Log .Trace ("\u0063\u0068\u006f\u0070\u0020\u0041\u0045\u0053\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u0020\u0028\u0025\u0064\u0029\u003a \u0025\u0020\u0078",len (buf ),buf );_fc .CryptBlocks (buf ,buf );_fb .Log .Trace ("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );if len (buf )==0{_fb .Log .Trace ("\u0045\u006d\u0070\u0074\u0079\u0020b\u0075\u0066\u002c\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067 \u0065\u006d\u0070\u0074\u0079\u0020\u0073t\u0072\u0069\u006e\u0067");return buf ,nil ;};_bc :=int (buf [len (buf )-1]);if _bc > len (buf ){_fb .Log .Debug ("\u0049\u006c\u006c\u0065g\u0061\u006c\u0020\u0070\u0061\u0064\u0020\u006c\u0065\u006eg\u0074h\u0020\u0028\u0025\u0064\u0020\u003e\u0020%\u0064\u0029",_bc ,len (buf ));return buf ,_ff .Errorf ("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0070a\u0064\u0020l\u0065\u006e\u0067\u0074\u0068");};buf =buf [:len (buf )-_bc ];return buf ,nil ;};

// HandlerVersion implements Filter interface.
func (_ged filterV2 )HandlerVersion ()(V ,R int ){V ,R =2,3;return ;};

// MakeKey implements Filter interface.
func (filterAESV2 )MakeKey (objNum ,genNum uint32 ,ekey []byte )([]byte ,error ){return _ac (objNum ,genNum ,ekey ,true );};

// PDFVersion implements Filter interface.
func (filterAESV3 )PDFVersion ()[2]int {return [2]int {2,0}};type filterAESV3 struct{filterAES };func (filterIdentity )HandlerVersion ()(V ,R int ){return ;};

// KeyLength implements Filter interface.
func (_dgc filterV2 )KeyLength ()int {return _dgc ._efc };func init (){_gdg ("\u0041\u0045\u0053V\u0033",_bg )};var (_cg =make (map[string ]filterFunc ););func (filterIdentity )MakeKey (objNum ,genNum uint32 ,fkey []byte )([]byte ,error ){return fkey ,nil };

// NewFilterAESV3 creates an AES-based filter with a 256 bit key (AESV3).
func NewFilterAESV3 ()Filter {_a ,_ecb :=_bg (FilterDict {});if _ecb !=nil {_fb .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0033\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_ecb );return filterAESV3 {};};return _a ;};

// KeyLength implements Filter interface.
func (filterAESV3 )KeyLength ()int {return 256/8};type filterAESV2 struct{filterAES };

// Name implements Filter interface.
func (filterAESV3 )Name ()string {return "\u0041\u0045\u0053V\u0033"};type filterFunc func (_gc FilterDict )(Filter ,error );func _dg (_ba FilterDict )(Filter ,error ){if _ba .Length %8!=0{return nil ,_ff .Errorf ("\u0063\u0072\u0079p\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020o\u0066\u0020\u0038\u0020\u0028\u0025\u0064\u0029",_ba .Length );};if _ba .Length < 5||_ba .Length > 16{if _ba .Length ==40||_ba .Length ==64||_ba .Length ==128{_fb .Log .Debug ("\u0053\u0054\u0041\u004e\u0044AR\u0044\u0020V\u0049\u004f\u004c\u0041\u0054\u0049\u004f\u004e\u003a\u0020\u0043\u0072\u0079\u0070\u0074\u0020\u004c\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072s\u0020\u0074\u006f \u0062\u0065\u0020\u0069\u006e\u0020\u0062\u0069\u0074\u0073\u0020\u0072\u0061t\u0068\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0020-\u0020\u0061s\u0073u\u006d\u0069\u006e\u0067\u0020\u0062\u0069t\u0073\u0020\u0028\u0025\u0064\u0029",_ba .Length );_ba .Length /=8;}else {return nil ,_ff .Errorf ("\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074h\u0020\u006e\u006f\u0074\u0020\u0069\u006e \u0072\u0061\u006e\u0067\u0065\u0020\u0034\u0030\u0020\u002d\u00201\u0032\u0038\u0020\u0062\u0069\u0074\u0020\u0028\u0025\u0064\u0029",_ba .Length );};};return filterV2 {_efc :_ba .Length },nil ;};

// Name implements Filter interface.
func (filterV2 )Name ()string {return "\u0056\u0032"};

// DecryptBytes implements Filter interface.
func (filterV2 )DecryptBytes (buf []byte ,okey []byte )([]byte ,error ){_bab ,_bac :=_fg .NewCipher (okey );if _bac !=nil {return nil ,_bac ;};_fb .Log .Trace ("\u0052\u00434\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078",buf );_bab .XORKeyStream (buf ,buf );_fb .Log .Trace ("\u0074o\u003a\u0020\u0025\u0020\u0078",buf );return buf ,nil ;};func _ec (_ede FilterDict )(Filter ,error ){if _ede .Length ==128{_fb .Log .Debug ("\u0041\u0045S\u0056\u0032\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029",_ede .Length );_ede .Length /=8;};if _ede .Length !=0&&_ede .Length !=16{return nil ,_ff .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0032\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029",_ede .Length );};return filterAESV2 {},nil ;};

// MakeKey implements Filter interface.
func (_gb filterV2 )MakeKey (objNum ,genNum uint32 ,ekey []byte )([]byte ,error ){return _ac (objNum ,genNum ,ekey ,false );};func init (){_gdg ("\u0056\u0032",_dg )};

// Filter is a common interface for crypt filter methods.
type Filter interface{

// Name returns a name of the filter that should be used in CFM field of Encrypt dictionary.
Name ()string ;

// KeyLength returns a length of the encryption key in bytes.
KeyLength ()int ;

// PDFVersion reports the minimal version of PDF document that introduced this filter.
PDFVersion ()[2]int ;

// HandlerVersion reports V and R parameters that should be used for this filter.
HandlerVersion ()(V ,R int );

// MakeKey generates a object encryption key based on file encryption key and object numbers.
// Used only for legacy filters - AESV3 doesn't change the key for each object.
MakeKey (_adb ,_dfc uint32 ,_gaa []byte )([]byte ,error );

// EncryptBytes encrypts a buffer using object encryption key, as returned by MakeKey.
// Implementation may reuse a buffer and encrypt data in-place.
EncryptBytes (_gafc []byte ,_aa []byte )([]byte ,error );

// DecryptBytes decrypts a buffer using object encryption key, as returned by MakeKey.
// Implementation may reuse a buffer and decrypt data in-place.
DecryptBytes (_bef []byte ,_ccg []byte )([]byte ,error );};

// PDFVersion implements Filter interface.
func (_cba filterV2 )PDFVersion ()[2]int {return [2]int {}};func _ca (_ccc string )(filterFunc ,error ){_fcd :=_cg [_ccc ];if _fcd ==nil {return nil ,_ff .Errorf ("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u0072\u0079p\u0074 \u0066\u0069\u006c\u0074\u0065\u0072\u003a \u0025\u0071",_ccc );};return _fcd ,nil ;};func _gdg (_bgg string ,_ag filterFunc ){if _ ,_ea :=_cg [_bgg ];_ea {panic ("\u0061l\u0072e\u0061\u0064\u0079\u0020\u0072e\u0067\u0069s\u0074\u0065\u0072\u0065\u0064");};_cg [_bgg ]=_ag ;};func (filterIdentity )KeyLength ()int {return 0};

// HandlerVersion implements Filter interface.
func (filterAESV2 )HandlerVersion ()(V ,R int ){V ,R =4,4;return ;};func (filterIdentity )DecryptBytes (p []byte ,okey []byte )([]byte ,error ){return p ,nil };

// Name implements Filter interface.
func (filterAESV2 )Name ()string {return "\u0041\u0045\u0053V\u0032"};func (filterIdentity )PDFVersion ()[2]int {return [2]int {}};

// NewFilterV2 creates a RC4-based filter with a specified key length (in bytes).
func NewFilterV2 (length int )Filter {_ab ,_ad :=_dg (FilterDict {Length :length });if _ad !=nil {_fb .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020R\u0043\u0034\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_ad );return filterV2 {_efc :length };};return _ab ;};

// MakeKey implements Filter interface.
func (filterAESV3 )MakeKey (_ ,_ uint32 ,ekey []byte )([]byte ,error ){return ekey ,nil };

// FilterDict represents information from a CryptFilter dictionary.
type FilterDict struct{CFM string ;AuthEvent _ed .AuthEvent ;Length int ;};func _ac (_gaf ,_feb uint32 ,_fa []byte ,_af bool )([]byte ,error ){_fd :=make ([]byte ,len (_fa )+5);for _eeb :=0;_eeb < len (_fa );_eeb ++{_fd [_eeb ]=_fa [_eeb ];};for _afd :=0;_afd < 3;_afd ++{_be :=byte ((_gaf >>uint32 (8*_afd ))&0xff);_fd [_afd +len (_fa )]=_be ;};for _egf :=0;_egf < 2;_egf ++{_dae :=byte ((_feb >>uint32 (8*_egf ))&0xff);_fd [_egf +len (_fa )+3]=_dae ;};if _af {_fd =append (_fd ,0x73);_fd =append (_fd ,0x41);_fd =append (_fd ,0x6C);_fd =append (_fd ,0x54);};_bff :=_g .New ();_bff .Write (_fd );_df :=_bff .Sum (nil );if len (_fa )+5< 16{return _df [0:len (_fa )+5],nil ;};return _df ,nil ;};var _ Filter =filterV2 {};

// PDFVersion implements Filter interface.
func (filterAESV2 )PDFVersion ()[2]int {return [2]int {1,5}};var _ Filter =filterAESV2 {};

// NewIdentity creates an identity filter that bypasses all data without changes.
func NewIdentity ()Filter {return filterIdentity {}};func (filterAES )EncryptBytes (buf []byte ,okey []byte )([]byte ,error ){_dc ,_da :=_d .NewCipher (okey );if _da !=nil {return nil ,_da ;};_fb .Log .Trace ("A\u0045\u0053\u0020\u0045nc\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );const _bfa =_d .BlockSize ;_ecbe :=_bfa -len (buf )%_bfa ;for _dd :=0;_dd < _ecbe ;_dd ++{buf =append (buf ,byte (_ecbe ));};_fb .Log .Trace ("\u0050a\u0064d\u0065\u0064\u0020\u0074\u006f \u0025\u0064 \u0062\u0079\u0074\u0065\u0073",len (buf ));_ga :=make ([]byte ,_bfa +len (buf ));_eg :=_ga [:_bfa ];if _ ,_fgb :=_b .ReadFull (_c .Reader ,_eg );_fgb !=nil {return nil ,_fgb ;};_cc :=_e .NewCBCEncrypter (_dc ,_eg );_cc .CryptBlocks (_ga [_bfa :],buf );buf =_ga ;_fb .Log .Trace ("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );return buf ,nil ;};type filterV2 struct{_efc int };var _ Filter =filterAESV3 {};type filterAES struct{};func _bg (_fe FilterDict )(Filter ,error ){if _fe .Length ==256{_fb .Log .Debug ("\u0041\u0045S\u0056\u0033\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029",_fe .Length );_fe .Length /=8;};if _fe .Length !=0&&_fe .Length !=32{return nil ,_ff .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0033\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029",_fe .Length );};return filterAESV3 {},nil ;};

// KeyLength implements Filter interface.
func (filterAESV2 )KeyLength ()int {return 128/8};type filterIdentity struct{};

// HandlerVersion implements Filter interface.
func (filterAESV3 )HandlerVersion ()(V ,R int ){V ,R =5,6;return ;};

// EncryptBytes implements Filter interface.
func (filterV2 )EncryptBytes (buf []byte ,okey []byte )([]byte ,error ){_bd ,_cf :=_fg .NewCipher (okey );if _cf !=nil {return nil ,_cf ;};_fb .Log .Trace ("\u0052\u00434\u0020\u0045\u006ec\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078",buf );_bd .XORKeyStream (buf ,buf );_fb .Log .Trace ("\u0074o\u003a\u0020\u0025\u0020\u0078",buf );return buf ,nil ;};

// NewFilterAESV2 creates an AES-based filter with a 128 bit key (AESV2).
func NewFilterAESV2 ()Filter {_ef ,_ce :=_ec (FilterDict {});if _ce !=nil {_fb .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_ce );return filterAESV2 {};};return _ef ;};func (filterIdentity )Name ()string {return "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"};