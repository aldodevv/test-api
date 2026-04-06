package constants

// RC (Response Code) Constants
// Menampung kode-kode balikan secara global agar konsisten
const (
	RCSukses       = "SC00"
	RCInvalidInput = "REQ99"
	RCDBError      = "DB99"
	RCUserError    = "USR99"
	RCAuthError    = "AUTH99"
)

// Response Message Constants
// Menampung pesan-pesan standar yang sering diulang
const (
	MsgSuksesAmbilData = "Berhasil mendapatkan data"
	MsgSuksesBikinData = "Berhasil membuat data"
	
	MsgInvalidInput    = "Format input tidak valid"
	MsgGagalSimpanUser = "Gagal menyimpan data user"
	MsgGagalAmbilUser  = "Gagal mengambil data user"
	MsgAksesDitolak    = "Akses ditolak! Token tidak valid/hilang."
)
