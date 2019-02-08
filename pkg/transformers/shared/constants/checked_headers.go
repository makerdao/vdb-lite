package constants

type TransformerExecution bool

const (
	HeaderRecheck    TransformerExecution = true
	HeaderMissing    TransformerExecution = false
	RecheckHeaderCap                      = "4"
)

var (
	BiteChecked               = "bite_checked"
	CatFileChopLumpChecked    = "cat_file_chop_lump_checked"
	CatFileFlipChecked        = "cat_file_flip_checked"
	CatFilePitVowChecked      = "cat_file_pit_vow_checked"
	DealChecked               = "deal_checked"
	DentChecked               = "dent_checked"
	DripDripChecked           = "drip_drip_checked"
	DripFileIlkChecked        = "drip_file_ilk_checked"
	DripFileRepoChecked       = "drip_file_repo_checked"
	DripFileVowChecked        = "drip_file_vow_checked"
	FlapKickChecked           = "flap_kick_checked"
	FlipKickChecked           = "flip_kick_checked"
	FlopKickChecked           = "flop_kick_checked"
	FrobChecked               = "frob_checked"
	PitFileDebtCeilingChecked = "pit_file_debt_ceiling_checked"
	PitFileIlkChecked         = "pit_file_ilk_checked"
	PriceFeedsChecked         = "price_feeds_checked"
	TendChecked               = "tend_checked"
	VatFluxChecked            = "vat_flux_checked"
	VatFoldChecked            = "vat_fold_checked"
	VatGrabChecked            = "vat_grab_checked"
	VatHealChecked            = "vat_heal_checked"
	VatInitChecked            = "vat_init_checked"
	VatMoveChecked            = "vat_move_checked"
	VatSlipChecked            = "vat_slip_checked"
	VatTollChecked            = "vat_toll_checked"
	VatTuneChecked            = "vat_tune_checked"
	VowFlogChecked            = "vow_flog_checked"
)
