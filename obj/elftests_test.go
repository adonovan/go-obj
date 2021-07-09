// Code generated by build.go. DO NOT EDIT.

package obj

import "github.com/aclements/go-obj/arch"

var elfTests = []*elfTest{
	{
		path: "hello-gcc10.3.0-I386-static",
		arch: arch.I386,
		sections: []Section{
			{Name: ".note.gnu.build-id", ID: 0, RawID: 1, Addr: 0x8048154, Size: 0x24},
			{Name: ".note.gnu.property", ID: 1, RawID: 2, Addr: 0x8048178, Size: 0x1c},
			{Name: ".note.ABI-tag", ID: 2, RawID: 3, Addr: 0x8048194, Size: 0x20},
			{Name: ".rel.plt", ID: 3, RawID: 4, Addr: 0x80481b4, Size: 0x70},
			{Name: ".init", ID: 4, RawID: 5, Addr: 0x8049000, Size: 0x24},
			{Name: ".plt", ID: 5, RawID: 6, Addr: 0x8049030, Size: 0xe0},
			{Name: ".text", ID: 6, RawID: 7, Addr: 0x8049110, Size: 0x6a7f1},
			{Name: "__libc_freeres_fn", ID: 7, RawID: 8, Addr: 0x80b3910, Size: 0xb55},
			{Name: ".fini", ID: 8, RawID: 9, Addr: 0x80b4468, Size: 0x18},
			{Name: ".rodata", ID: 9, RawID: 10, Addr: 0x80b5000, Size: 0x1c3d0},
			{Name: ".eh_frame", ID: 10, RawID: 11, Addr: 0x80d13d0, Size: 0x13f74},
			{Name: ".gcc_except_table", ID: 11, RawID: 12, Addr: 0x80e5344, Size: 0xc4},
			{Name: ".tdata", ID: 12, RawID: 13, Addr: 0x80e65e0, Size: 0x10},
			{Name: ".tbss", ID: 13, RawID: 14, Addr: 0x80e65f0, Size: 0x20},
			{Name: ".init_array", ID: 14, RawID: 15, Addr: 0x80e65f0, Size: 0x8},
			{Name: ".fini_array", ID: 15, RawID: 16, Addr: 0x80e65f8, Size: 0x8},
			{Name: ".data.rel.ro", ID: 16, RawID: 17, Addr: 0x80e6600, Size: 0x19d4},
			{Name: ".got", ID: 17, RawID: 18, Addr: 0x80e7fd4, Size: 0x24},
			{Name: ".got.plt", ID: 18, RawID: 19, Addr: 0x80e8000, Size: 0x44},
			{Name: ".data", ID: 19, RawID: 20, Addr: 0x80e8060, Size: 0xec0},
			{Name: "__libc_subfreeres", ID: 20, RawID: 21, Addr: 0x80e8f20, Size: 0x24},
			{Name: "__libc_IO_vtables", ID: 21, RawID: 22, Addr: 0x80e8f60, Size: 0x3b4},
			{Name: "__libc_atexit", ID: 22, RawID: 23, Addr: 0x80e9314, Size: 0x4},
			{Name: ".bss", ID: 23, RawID: 24, Addr: 0x80e9320, Size: 0xda4},
			{Name: "__libc_freeres_ptrs", ID: 24, RawID: 25, Addr: 0x80ea0c4, Size: 0x10},
			{Name: ".comment", ID: 25, RawID: 26, Addr: 0x0, Size: 0x2b},
			{Name: ".debug_aranges", ID: 26, RawID: 27, Addr: 0x0, Size: 0x20},
			{Name: ".debug_info", ID: 27, RawID: 28, Addr: 0x0, Size: 0xb8},
			{Name: ".debug_abbrev", ID: 28, RawID: 29, Addr: 0x0, Size: 0x62},
			{Name: ".debug_line", ID: 29, RawID: 30, Addr: 0x0, Size: 0x42},
			{Name: ".debug_str", ID: 30, RawID: 31, Addr: 0x0, Size: 0x10c},
			{Name: ".symtab", ID: 31, RawID: 32, Addr: 0x0, Size: 0x8b90},
			{Name: ".strtab", ID: 32, RawID: 33, Addr: 0x0, Size: 0x67fc},
			{Name: ".shstrtab", ID: 33, RawID: 34, Addr: 0x0, Size: 0x17a},
		},
		nSyms: 2232,
	},
	{
		path: "hello-gcc10.3.0-I386-dyn",
		arch: arch.I386,
		sections: []Section{
			{Name: ".interp", ID: 0, RawID: 1, Addr: 0x80481b4, Size: 0x13},
			{Name: ".note.gnu.build-id", ID: 1, RawID: 2, Addr: 0x80481c8, Size: 0x24},
			{Name: ".note.gnu.property", ID: 2, RawID: 3, Addr: 0x80481ec, Size: 0x1c},
			{Name: ".note.ABI-tag", ID: 3, RawID: 4, Addr: 0x8048208, Size: 0x20},
			{Name: ".gnu.hash", ID: 4, RawID: 5, Addr: 0x8048228, Size: 0x20},
			{Name: ".dynsym", ID: 5, RawID: 6, Addr: 0x8048248, Size: 0x50},
			{Name: ".dynstr", ID: 6, RawID: 7, Addr: 0x8048298, Size: 0x4a},
			{Name: ".gnu.version", ID: 7, RawID: 8, Addr: 0x80482e2, Size: 0xa},
			{Name: ".gnu.version_r", ID: 8, RawID: 9, Addr: 0x80482ec, Size: 0x20},
			{Name: ".rel.dyn", ID: 9, RawID: 10, Addr: 0x804830c, Size: 0x8},
			{Name: ".rel.plt", ID: 10, RawID: 11, Addr: 0x8048314, Size: 0x10},
			{Name: ".init", ID: 11, RawID: 12, Addr: 0x8049000, Size: 0x24},
			{Name: ".plt", ID: 12, RawID: 13, Addr: 0x8049030, Size: 0x30},
			{Name: ".plt.sec", ID: 13, RawID: 14, Addr: 0x8049060, Size: 0x20},
			{Name: ".text", ID: 14, RawID: 15, Addr: 0x8049080, Size: 0x1d9},
			{Name: ".fini", ID: 15, RawID: 16, Addr: 0x804925c, Size: 0x18},
			{Name: ".rodata", ID: 16, RawID: 17, Addr: 0x804a000, Size: 0xe},
			{Name: ".eh_frame_hdr", ID: 17, RawID: 18, Addr: 0x804a010, Size: 0x54},
			{Name: ".eh_frame", ID: 18, RawID: 19, Addr: 0x804a064, Size: 0x138},
			{Name: ".init_array", ID: 19, RawID: 20, Addr: 0x804bf0c, Size: 0x4},
			{Name: ".fini_array", ID: 20, RawID: 21, Addr: 0x804bf10, Size: 0x4},
			{Name: ".dynamic", ID: 21, RawID: 22, Addr: 0x804bf14, Size: 0xe8},
			{Name: ".got", ID: 22, RawID: 23, Addr: 0x804bffc, Size: 0x4},
			{Name: ".got.plt", ID: 23, RawID: 24, Addr: 0x804c000, Size: 0x14},
			{Name: ".data", ID: 24, RawID: 25, Addr: 0x804c014, Size: 0x8},
			{Name: ".bss", ID: 25, RawID: 26, Addr: 0x804c01c, Size: 0x4},
			{Name: ".comment", ID: 26, RawID: 27, Addr: 0x0, Size: 0x2b},
			{Name: ".debug_aranges", ID: 27, RawID: 28, Addr: 0x0, Size: 0x20},
			{Name: ".debug_info", ID: 28, RawID: 29, Addr: 0x0, Size: 0xb8},
			{Name: ".debug_abbrev", ID: 29, RawID: 30, Addr: 0x0, Size: 0x62},
			{Name: ".debug_line", ID: 30, RawID: 31, Addr: 0x0, Size: 0x42},
			{Name: ".debug_str", ID: 31, RawID: 32, Addr: 0x0, Size: 0x10c},
			{Name: ".symtab", ID: 32, RawID: 33, Addr: 0x0, Size: 0x4a0},
			{Name: ".strtab", ID: 33, RawID: 34, Addr: 0x0, Size: 0x24f},
			{Name: ".shstrtab", ID: 34, RawID: 35, Addr: 0x0, Size: 0x15d},
		},
		nSyms:    77,
		textData: parseHex(`f30f1efb31ed5e89e183e4f0505452e82300000081c36c2f00008d8350d2ffff508d83e0d1ffff505156c7c09691040850e8bafffffff48b1c24c36690669090f30f1efbc366906690669066906690908b1c24c3669066906690669066906690b81cc004083d1cc004087424b80000000085c0741b5589e583ec14681cc00408ffd083c410c9c38db426000000006690c38db426000000008db4260000000090b81cc004082d1cc0040889c2c1e81fc1fa0201d0d1f87420ba0000000085d274175589e583ec1050681cc00408ffd283c410c9c38d742600c38db42600000000f30f1efb803d1cc0040800751b5589e583ec08e868ffffffc6051cc0040801c9c38db42600000000c38db42600000000f30f1efbeb8af30f1efb8d4c240483e4f0ff71fc5589e55351e82800000005522e000083ec0c8d9008e0ffff5289c3e89cfeffff83c410b8000000008d65f8595b5d8d61fcc38b0424c3669066906690f30f1efb55e86b00000081c5162e000057565383ec0c89eb8b7c2428e8fffdffff8d9d10ffffff8d850cffffff29c3c1fb02742931f68db426000000008d760083ec0457ff74242cff74242cff94b50cffffff83c60183c41039f375e383c40c5b5e5f5dc38db426000000008d742600f30f1efbc38b2c24c3`),
	},
	{
		path: "hello-gcc10.3.0-I386-dyn-stripped",
		arch: arch.I386,
		sections: []Section{
			{Name: ".interp", ID: 0, RawID: 1, Addr: 0x80481b4, Size: 0x13},
			{Name: ".note.gnu.build-id", ID: 1, RawID: 2, Addr: 0x80481c8, Size: 0x24},
			{Name: ".note.gnu.property", ID: 2, RawID: 3, Addr: 0x80481ec, Size: 0x1c},
			{Name: ".note.ABI-tag", ID: 3, RawID: 4, Addr: 0x8048208, Size: 0x20},
			{Name: ".gnu.hash", ID: 4, RawID: 5, Addr: 0x8048228, Size: 0x20},
			{Name: ".dynsym", ID: 5, RawID: 6, Addr: 0x8048248, Size: 0x50},
			{Name: ".dynstr", ID: 6, RawID: 7, Addr: 0x8048298, Size: 0x4a},
			{Name: ".gnu.version", ID: 7, RawID: 8, Addr: 0x80482e2, Size: 0xa},
			{Name: ".gnu.version_r", ID: 8, RawID: 9, Addr: 0x80482ec, Size: 0x20},
			{Name: ".rel.dyn", ID: 9, RawID: 10, Addr: 0x804830c, Size: 0x8},
			{Name: ".rel.plt", ID: 10, RawID: 11, Addr: 0x8048314, Size: 0x10},
			{Name: ".init", ID: 11, RawID: 12, Addr: 0x8049000, Size: 0x24},
			{Name: ".plt", ID: 12, RawID: 13, Addr: 0x8049030, Size: 0x30},
			{Name: ".plt.sec", ID: 13, RawID: 14, Addr: 0x8049060, Size: 0x20},
			{Name: ".text", ID: 14, RawID: 15, Addr: 0x8049080, Size: 0x1d9},
			{Name: ".fini", ID: 15, RawID: 16, Addr: 0x804925c, Size: 0x18},
			{Name: ".rodata", ID: 16, RawID: 17, Addr: 0x804a000, Size: 0xe},
			{Name: ".eh_frame_hdr", ID: 17, RawID: 18, Addr: 0x804a010, Size: 0x54},
			{Name: ".eh_frame", ID: 18, RawID: 19, Addr: 0x804a064, Size: 0x138},
			{Name: ".init_array", ID: 19, RawID: 20, Addr: 0x804bf0c, Size: 0x4},
			{Name: ".fini_array", ID: 20, RawID: 21, Addr: 0x804bf10, Size: 0x4},
			{Name: ".dynamic", ID: 21, RawID: 22, Addr: 0x804bf14, Size: 0xe8},
			{Name: ".got", ID: 22, RawID: 23, Addr: 0x804bffc, Size: 0x4},
			{Name: ".got.plt", ID: 23, RawID: 24, Addr: 0x804c000, Size: 0x14},
			{Name: ".data", ID: 24, RawID: 25, Addr: 0x804c014, Size: 0x8},
			{Name: ".bss", ID: 25, RawID: 26, Addr: 0x804c01c, Size: 0x4},
			{Name: ".comment", ID: 26, RawID: 27, Addr: 0x0, Size: 0x2b},
			{Name: ".shstrtab", ID: 27, RawID: 28, Addr: 0x0, Size: 0x10d},
		},
		nSyms: 4,
	},
	{
		path: "hello-gcc10.3.0-I386-pie",
		arch: arch.I386,
		sections: []Section{
			{Name: ".interp", ID: 0, RawID: 1, Addr: 0x1b4, Size: 0x13},
			{Name: ".note.gnu.build-id", ID: 1, RawID: 2, Addr: 0x1c8, Size: 0x24},
			{Name: ".note.gnu.property", ID: 2, RawID: 3, Addr: 0x1ec, Size: 0x1c},
			{Name: ".note.ABI-tag", ID: 3, RawID: 4, Addr: 0x208, Size: 0x20},
			{Name: ".gnu.hash", ID: 4, RawID: 5, Addr: 0x228, Size: 0x20},
			{Name: ".dynsym", ID: 5, RawID: 6, Addr: 0x248, Size: 0x80},
			{Name: ".dynstr", ID: 6, RawID: 7, Addr: 0x2c8, Size: 0x9b},
			{Name: ".gnu.version", ID: 7, RawID: 8, Addr: 0x364, Size: 0x10},
			{Name: ".gnu.version_r", ID: 8, RawID: 9, Addr: 0x374, Size: 0x30},
			{Name: ".rel.dyn", ID: 9, RawID: 10, Addr: 0x3a4, Size: 0x40},
			{Name: ".rel.plt", ID: 10, RawID: 11, Addr: 0x3e4, Size: 0x10},
			{Name: ".init", ID: 11, RawID: 12, Addr: 0x1000, Size: 0x24},
			{Name: ".plt", ID: 12, RawID: 13, Addr: 0x1030, Size: 0x30},
			{Name: ".plt.got", ID: 13, RawID: 14, Addr: 0x1060, Size: 0x10},
			{Name: ".plt.sec", ID: 14, RawID: 15, Addr: 0x1070, Size: 0x20},
			{Name: ".text", ID: 15, RawID: 16, Addr: 0x1090, Size: 0x209},
			{Name: ".fini", ID: 16, RawID: 17, Addr: 0x129c, Size: 0x18},
			{Name: ".rodata", ID: 17, RawID: 18, Addr: 0x2000, Size: 0xe},
			{Name: ".eh_frame_hdr", ID: 18, RawID: 19, Addr: 0x2010, Size: 0x54},
			{Name: ".eh_frame", ID: 19, RawID: 20, Addr: 0x2064, Size: 0x138},
			{Name: ".init_array", ID: 20, RawID: 21, Addr: 0x3ed8, Size: 0x4},
			{Name: ".fini_array", ID: 21, RawID: 22, Addr: 0x3edc, Size: 0x4},
			{Name: ".dynamic", ID: 22, RawID: 23, Addr: 0x3ee0, Size: 0xf8},
			{Name: ".got", ID: 23, RawID: 24, Addr: 0x3fd8, Size: 0x28},
			{Name: ".data", ID: 24, RawID: 25, Addr: 0x4000, Size: 0x8},
			{Name: ".bss", ID: 25, RawID: 26, Addr: 0x4008, Size: 0x4},
			{Name: ".comment", ID: 26, RawID: 27, Addr: 0x0, Size: 0x2b},
			{Name: ".debug_aranges", ID: 27, RawID: 28, Addr: 0x0, Size: 0x20},
			{Name: ".debug_info", ID: 28, RawID: 29, Addr: 0x0, Size: 0xb8},
			{Name: ".debug_abbrev", ID: 29, RawID: 30, Addr: 0x0, Size: 0x62},
			{Name: ".debug_line", ID: 30, RawID: 31, Addr: 0x0, Size: 0x42},
			{Name: ".debug_str", ID: 31, RawID: 32, Addr: 0x0, Size: 0x10c},
			{Name: ".symtab", ID: 32, RawID: 33, Addr: 0x0, Size: 0x4d0},
			{Name: ".strtab", ID: 33, RawID: 34, Addr: 0x0, Size: 0x2a0},
			{Name: ".shstrtab", ID: 34, RawID: 35, Addr: 0x0, Size: 0x158},
		},
		nSyms: 83,
	},
	{
		path: "hello-gcc10.3.0-I386-rel.o",
		arch: arch.I386,
		sections: []Section{
			{Name: ".group", ID: 0, RawID: 1, Addr: 0x0, Size: 0x8},
			{Name: ".text", ID: 1, RawID: 2, Addr: 0x0, Size: 0x40},
			{Name: ".rel.text", ID: 2, RawID: 3, Addr: 0x0, Size: 0x20},
			{Name: ".data", ID: 3, RawID: 4, Addr: 0x0, Size: 0x0},
			{Name: ".bss", ID: 4, RawID: 5, Addr: 0x0, Size: 0x0},
			{Name: ".rodata", ID: 5, RawID: 6, Addr: 0x0, Size: 0x6},
			{Name: ".text.__x86.get_pc_thunk.ax", ID: 6, RawID: 7, Addr: 0x0, Size: 0x4},
			{Name: ".debug_info", ID: 7, RawID: 8, Addr: 0x0, Size: 0xb8},
			{Name: ".rel.debug_info", ID: 8, RawID: 9, Addr: 0x0, Size: 0xa0},
			{Name: ".debug_abbrev", ID: 9, RawID: 10, Addr: 0x0, Size: 0x62},
			{Name: ".debug_aranges", ID: 10, RawID: 11, Addr: 0x0, Size: 0x20},
			{Name: ".rel.debug_aranges", ID: 11, RawID: 12, Addr: 0x0, Size: 0x10},
			{Name: ".debug_line", ID: 12, RawID: 13, Addr: 0x0, Size: 0x42},
			{Name: ".rel.debug_line", ID: 13, RawID: 14, Addr: 0x0, Size: 0x8},
			{Name: ".debug_str", ID: 14, RawID: 15, Addr: 0x0, Size: 0x145},
			{Name: ".comment", ID: 15, RawID: 16, Addr: 0x0, Size: 0x2c},
			{Name: ".note.GNU-stack", ID: 16, RawID: 17, Addr: 0x0, Size: 0x0},
			{Name: ".note.gnu.property", ID: 17, RawID: 18, Addr: 0x0, Size: 0x1c},
			{Name: ".eh_frame", ID: 18, RawID: 19, Addr: 0x0, Size: 0x60},
			{Name: ".rel.eh_frame", ID: 19, RawID: 20, Addr: 0x0, Size: 0x10},
			{Name: ".symtab", ID: 20, RawID: 21, Addr: 0x0, Size: 0x150},
			{Name: ".strtab", ID: 21, RawID: 22, Addr: 0x0, Size: 0x3f},
			{Name: ".shstrtab", ID: 22, RawID: 23, Addr: 0x0, Size: 0xe1},
		},
		nSyms:    20,
		textData: parseHex(`f30f1efb8d4c240483e4f0ff71fc5589e55351e8fcffffff050100000083ec0c8d90000000005289c3e8fcffffff83c410b8000000008d65f8595b5d8d61fcc3`),
	},
	{
		path: "hello-gcc10.3.0-AMD64-static",
		arch: arch.AMD64,
		sections: []Section{
			{Name: ".note.gnu.property", ID: 0, RawID: 1, Addr: 0x400270, Size: 0x20},
			{Name: ".note.gnu.build-id", ID: 1, RawID: 2, Addr: 0x400290, Size: 0x24},
			{Name: ".note.ABI-tag", ID: 2, RawID: 3, Addr: 0x4002b4, Size: 0x20},
			{Name: ".rela.plt", ID: 3, RawID: 4, Addr: 0x4002d8, Size: 0x240},
			{Name: ".init", ID: 4, RawID: 5, Addr: 0x401000, Size: 0x1b},
			{Name: ".plt", ID: 5, RawID: 6, Addr: 0x401020, Size: 0x180},
			{Name: ".text", ID: 6, RawID: 7, Addr: 0x4011a0, Size: 0x8a470},
			{Name: "__libc_freeres_fn", ID: 7, RawID: 8, Addr: 0x48b610, Size: 0x14f0},
			{Name: ".fini", ID: 8, RawID: 9, Addr: 0x48cb00, Size: 0xd},
			{Name: ".rodata", ID: 9, RawID: 10, Addr: 0x48d000, Size: 0x1d06c},
			{Name: ".stapsdt.base", ID: 10, RawID: 11, Addr: 0x4aa06c, Size: 0x1},
			{Name: ".eh_frame", ID: 11, RawID: 12, Addr: 0x4aa070, Size: 0xac44},
			{Name: ".gcc_except_table", ID: 12, RawID: 13, Addr: 0x4b4cb4, Size: 0xc4},
			{Name: ".tdata", ID: 13, RawID: 14, Addr: 0x4b5fe0, Size: 0x20},
			{Name: ".tbss", ID: 14, RawID: 15, Addr: 0x4b6000, Size: 0x40},
			{Name: ".init_array", ID: 15, RawID: 16, Addr: 0x4b6000, Size: 0x10},
			{Name: ".fini_array", ID: 16, RawID: 17, Addr: 0x4b6010, Size: 0x10},
			{Name: ".data.rel.ro", ID: 17, RawID: 18, Addr: 0x4b6020, Size: 0x2ed4},
			{Name: ".got", ID: 18, RawID: 19, Addr: 0x4b8ef8, Size: 0xf0},
			{Name: ".got.plt", ID: 19, RawID: 20, Addr: 0x4b9000, Size: 0xd8},
			{Name: ".data", ID: 20, RawID: 21, Addr: 0x4b90e0, Size: 0x1a70},
			{Name: "__libc_subfreeres", ID: 21, RawID: 22, Addr: 0x4bab50, Size: 0x48},
			{Name: "__libc_IO_vtables", ID: 22, RawID: 23, Addr: 0x4baba0, Size: 0x768},
			{Name: "__libc_atexit", ID: 23, RawID: 24, Addr: 0x4bb308, Size: 0x8},
			{Name: ".bss", ID: 24, RawID: 25, Addr: 0x4bb320, Size: 0x1800},
			{Name: "__libc_freeres_ptrs", ID: 25, RawID: 26, Addr: 0x4bcb20, Size: 0x20},
			{Name: ".comment", ID: 26, RawID: 27, Addr: 0x0, Size: 0x2b},
			{Name: ".note.stapsdt", ID: 27, RawID: 28, Addr: 0x0, Size: 0x1390},
			{Name: ".debug_aranges", ID: 28, RawID: 29, Addr: 0x0, Size: 0x30},
			{Name: ".debug_info", ID: 29, RawID: 30, Addr: 0x0, Size: 0xba},
			{Name: ".debug_abbrev", ID: 30, RawID: 31, Addr: 0x0, Size: 0x62},
			{Name: ".debug_line", ID: 31, RawID: 32, Addr: 0x0, Size: 0x45},
			{Name: ".debug_str", ID: 32, RawID: 33, Addr: 0x0, Size: 0x104},
			{Name: ".symtab", ID: 33, RawID: 34, Addr: 0x0, Size: 0xb2e0},
			{Name: ".strtab", ID: 34, RawID: 35, Addr: 0x0, Size: 0x685f},
			{Name: ".shstrtab", ID: 35, RawID: 36, Addr: 0x0, Size: 0x197},
		},
		nSyms: 1907,
	},
	{
		path: "hello-gcc10.3.0-AMD64-dyn",
		arch: arch.AMD64,
		sections: []Section{
			{Name: ".interp", ID: 0, RawID: 1, Addr: 0x400318, Size: 0x1c},
			{Name: ".note.gnu.property", ID: 1, RawID: 2, Addr: 0x400338, Size: 0x20},
			{Name: ".note.gnu.build-id", ID: 2, RawID: 3, Addr: 0x400358, Size: 0x24},
			{Name: ".note.ABI-tag", ID: 3, RawID: 4, Addr: 0x40037c, Size: 0x20},
			{Name: ".gnu.hash", ID: 4, RawID: 5, Addr: 0x4003a0, Size: 0x1c},
			{Name: ".dynsym", ID: 5, RawID: 6, Addr: 0x4003c0, Size: 0x60},
			{Name: ".dynstr", ID: 6, RawID: 7, Addr: 0x400420, Size: 0x3d},
			{Name: ".gnu.version", ID: 7, RawID: 8, Addr: 0x40045e, Size: 0x8},
			{Name: ".gnu.version_r", ID: 8, RawID: 9, Addr: 0x400468, Size: 0x20},
			{Name: ".rela.dyn", ID: 9, RawID: 10, Addr: 0x400488, Size: 0x30},
			{Name: ".rela.plt", ID: 10, RawID: 11, Addr: 0x4004b8, Size: 0x18},
			{Name: ".init", ID: 11, RawID: 12, Addr: 0x401000, Size: 0x1b},
			{Name: ".plt", ID: 12, RawID: 13, Addr: 0x401020, Size: 0x20},
			{Name: ".plt.sec", ID: 13, RawID: 14, Addr: 0x401040, Size: 0x10},
			{Name: ".text", ID: 14, RawID: 15, Addr: 0x401050, Size: 0x185},
			{Name: ".fini", ID: 15, RawID: 16, Addr: 0x4011d8, Size: 0xd},
			{Name: ".rodata", ID: 16, RawID: 17, Addr: 0x402000, Size: 0xa},
			{Name: ".eh_frame_hdr", ID: 17, RawID: 18, Addr: 0x40200c, Size: 0x44},
			{Name: ".eh_frame", ID: 18, RawID: 19, Addr: 0x402050, Size: 0x100},
			{Name: ".init_array", ID: 19, RawID: 20, Addr: 0x403e10, Size: 0x8},
			{Name: ".fini_array", ID: 20, RawID: 21, Addr: 0x403e18, Size: 0x8},
			{Name: ".dynamic", ID: 21, RawID: 22, Addr: 0x403e20, Size: 0x1d0},
			{Name: ".got", ID: 22, RawID: 23, Addr: 0x403ff0, Size: 0x10},
			{Name: ".got.plt", ID: 23, RawID: 24, Addr: 0x404000, Size: 0x20},
			{Name: ".data", ID: 24, RawID: 25, Addr: 0x404020, Size: 0x10},
			{Name: ".bss", ID: 25, RawID: 26, Addr: 0x404030, Size: 0x8},
			{Name: ".comment", ID: 26, RawID: 27, Addr: 0x0, Size: 0x2b},
			{Name: ".debug_aranges", ID: 27, RawID: 28, Addr: 0x0, Size: 0x30},
			{Name: ".debug_info", ID: 28, RawID: 29, Addr: 0x0, Size: 0xba},
			{Name: ".debug_abbrev", ID: 29, RawID: 30, Addr: 0x0, Size: 0x62},
			{Name: ".debug_line", ID: 30, RawID: 31, Addr: 0x0, Size: 0x45},
			{Name: ".debug_str", ID: 31, RawID: 32, Addr: 0x0, Size: 0x104},
			{Name: ".symtab", ID: 32, RawID: 33, Addr: 0x0, Size: 0x690},
			{Name: ".strtab", ID: 33, RawID: 34, Addr: 0x0, Size: 0x212},
			{Name: ".shstrtab", ID: 34, RawID: 35, Addr: 0x0, Size: 0x15f},
		},
		nSyms:    72,
		textData: parseHex(`f30f1efa31ed4989d15e4889e24883e4f0505449c7c0d011400048c7c16011400048c7c736114000ff15722f0000f490f30f1efac3662e0f1f84000000000090b830404000483d304040007413b8000000004885c07409bf30404000ffe06690c366662e0f1f8400000000000f1f4000be304040004881ee304040004889f048c1ee3f48c1f8034801c648d1fe7411b8000000004885c07407bf30404000ffe0c366662e0f1f8400000000000f1f4000f30f1efa803d252f0000007513554889e5e87affffffc605132f0000015dc390c366662e0f1f8400000000000f1f4000f30f1efaeb8af30f1efa554889e54883ec10897dfc488975f0488d3db40e0000e8ebfeffffb800000000c9c30f1f4000f30f1efa41574c8d3da32c000041564989d641554989f541544189fc55488d2d942c0000534c29fd4883ec08e86ffeffff48c1fd03741f31db0f1f80000000004c89f24c89ee4489e741ff14df4883c3014839dd75ea4883c4085b5d415c415d415e415fc366662e0f1f840000000000f30f1efac3`),
	},
	{
		path: "hello-gcc10.3.0-AMD64-dyn-stripped",
		arch: arch.AMD64,
		sections: []Section{
			{Name: ".interp", ID: 0, RawID: 1, Addr: 0x400318, Size: 0x1c},
			{Name: ".note.gnu.property", ID: 1, RawID: 2, Addr: 0x400338, Size: 0x20},
			{Name: ".note.gnu.build-id", ID: 2, RawID: 3, Addr: 0x400358, Size: 0x24},
			{Name: ".note.ABI-tag", ID: 3, RawID: 4, Addr: 0x40037c, Size: 0x20},
			{Name: ".gnu.hash", ID: 4, RawID: 5, Addr: 0x4003a0, Size: 0x1c},
			{Name: ".dynsym", ID: 5, RawID: 6, Addr: 0x4003c0, Size: 0x60},
			{Name: ".dynstr", ID: 6, RawID: 7, Addr: 0x400420, Size: 0x3d},
			{Name: ".gnu.version", ID: 7, RawID: 8, Addr: 0x40045e, Size: 0x8},
			{Name: ".gnu.version_r", ID: 8, RawID: 9, Addr: 0x400468, Size: 0x20},
			{Name: ".rela.dyn", ID: 9, RawID: 10, Addr: 0x400488, Size: 0x30},
			{Name: ".rela.plt", ID: 10, RawID: 11, Addr: 0x4004b8, Size: 0x18},
			{Name: ".init", ID: 11, RawID: 12, Addr: 0x401000, Size: 0x1b},
			{Name: ".plt", ID: 12, RawID: 13, Addr: 0x401020, Size: 0x20},
			{Name: ".plt.sec", ID: 13, RawID: 14, Addr: 0x401040, Size: 0x10},
			{Name: ".text", ID: 14, RawID: 15, Addr: 0x401050, Size: 0x185},
			{Name: ".fini", ID: 15, RawID: 16, Addr: 0x4011d8, Size: 0xd},
			{Name: ".rodata", ID: 16, RawID: 17, Addr: 0x402000, Size: 0xa},
			{Name: ".eh_frame_hdr", ID: 17, RawID: 18, Addr: 0x40200c, Size: 0x44},
			{Name: ".eh_frame", ID: 18, RawID: 19, Addr: 0x402050, Size: 0x100},
			{Name: ".init_array", ID: 19, RawID: 20, Addr: 0x403e10, Size: 0x8},
			{Name: ".fini_array", ID: 20, RawID: 21, Addr: 0x403e18, Size: 0x8},
			{Name: ".dynamic", ID: 21, RawID: 22, Addr: 0x403e20, Size: 0x1d0},
			{Name: ".got", ID: 22, RawID: 23, Addr: 0x403ff0, Size: 0x10},
			{Name: ".got.plt", ID: 23, RawID: 24, Addr: 0x404000, Size: 0x20},
			{Name: ".data", ID: 24, RawID: 25, Addr: 0x404020, Size: 0x10},
			{Name: ".bss", ID: 25, RawID: 26, Addr: 0x404030, Size: 0x8},
			{Name: ".comment", ID: 26, RawID: 27, Addr: 0x0, Size: 0x2b},
			{Name: ".shstrtab", ID: 27, RawID: 28, Addr: 0x0, Size: 0x10f},
		},
		nSyms: 3,
	},
	{
		path: "hello-gcc10.3.0-AMD64-pie",
		arch: arch.AMD64,
		sections: []Section{
			{Name: ".interp", ID: 0, RawID: 1, Addr: 0x318, Size: 0x1c},
			{Name: ".note.gnu.property", ID: 1, RawID: 2, Addr: 0x338, Size: 0x20},
			{Name: ".note.gnu.build-id", ID: 2, RawID: 3, Addr: 0x358, Size: 0x24},
			{Name: ".note.ABI-tag", ID: 3, RawID: 4, Addr: 0x37c, Size: 0x20},
			{Name: ".gnu.hash", ID: 4, RawID: 5, Addr: 0x3a0, Size: 0x24},
			{Name: ".dynsym", ID: 5, RawID: 6, Addr: 0x3c8, Size: 0xa8},
			{Name: ".dynstr", ID: 6, RawID: 7, Addr: 0x470, Size: 0x82},
			{Name: ".gnu.version", ID: 7, RawID: 8, Addr: 0x4f2, Size: 0xe},
			{Name: ".gnu.version_r", ID: 8, RawID: 9, Addr: 0x500, Size: 0x20},
			{Name: ".rela.dyn", ID: 9, RawID: 10, Addr: 0x520, Size: 0xc0},
			{Name: ".rela.plt", ID: 10, RawID: 11, Addr: 0x5e0, Size: 0x18},
			{Name: ".init", ID: 11, RawID: 12, Addr: 0x1000, Size: 0x1b},
			{Name: ".plt", ID: 12, RawID: 13, Addr: 0x1020, Size: 0x20},
			{Name: ".plt.got", ID: 13, RawID: 14, Addr: 0x1040, Size: 0x10},
			{Name: ".plt.sec", ID: 14, RawID: 15, Addr: 0x1050, Size: 0x10},
			{Name: ".text", ID: 15, RawID: 16, Addr: 0x1060, Size: 0x185},
			{Name: ".fini", ID: 16, RawID: 17, Addr: 0x11e8, Size: 0xd},
			{Name: ".rodata", ID: 17, RawID: 18, Addr: 0x2000, Size: 0xa},
			{Name: ".eh_frame_hdr", ID: 18, RawID: 19, Addr: 0x200c, Size: 0x44},
			{Name: ".eh_frame", ID: 19, RawID: 20, Addr: 0x2050, Size: 0x108},
			{Name: ".init_array", ID: 20, RawID: 21, Addr: 0x3db8, Size: 0x8},
			{Name: ".fini_array", ID: 21, RawID: 22, Addr: 0x3dc0, Size: 0x8},
			{Name: ".dynamic", ID: 22, RawID: 23, Addr: 0x3dc8, Size: 0x1f0},
			{Name: ".got", ID: 23, RawID: 24, Addr: 0x3fb8, Size: 0x48},
			{Name: ".data", ID: 24, RawID: 25, Addr: 0x4000, Size: 0x10},
			{Name: ".bss", ID: 25, RawID: 26, Addr: 0x4010, Size: 0x8},
			{Name: ".comment", ID: 26, RawID: 27, Addr: 0x0, Size: 0x2b},
			{Name: ".debug_aranges", ID: 27, RawID: 28, Addr: 0x0, Size: 0x30},
			{Name: ".debug_info", ID: 28, RawID: 29, Addr: 0x0, Size: 0xba},
			{Name: ".debug_abbrev", ID: 29, RawID: 30, Addr: 0x0, Size: 0x62},
			{Name: ".debug_line", ID: 30, RawID: 31, Addr: 0x0, Size: 0x45},
			{Name: ".debug_str", ID: 31, RawID: 32, Addr: 0x0, Size: 0x104},
			{Name: ".symtab", ID: 32, RawID: 33, Addr: 0x0, Size: 0x6c0},
			{Name: ".strtab", ID: 33, RawID: 34, Addr: 0x0, Size: 0x24d},
			{Name: ".shstrtab", ID: 34, RawID: 35, Addr: 0x0, Size: 0x15a},
		},
		nSyms: 77,
	},
	{
		path: "hello-gcc10.3.0-AMD64-rel.o",
		arch: arch.AMD64,
		sections: []Section{
			{Name: ".text", ID: 0, RawID: 1, Addr: 0x0, Size: 0x26},
			{Name: ".rela.text", ID: 1, RawID: 2, Addr: 0x0, Size: 0x30},
			{Name: ".data", ID: 2, RawID: 3, Addr: 0x0, Size: 0x0},
			{Name: ".bss", ID: 3, RawID: 4, Addr: 0x0, Size: 0x0},
			{Name: ".rodata", ID: 4, RawID: 5, Addr: 0x0, Size: 0x6},
			{Name: ".debug_info", ID: 5, RawID: 6, Addr: 0x0, Size: 0xba},
			{Name: ".rela.debug_info", ID: 6, RawID: 7, Addr: 0x0, Size: 0x1b0},
			{Name: ".debug_abbrev", ID: 7, RawID: 8, Addr: 0x0, Size: 0x62},
			{Name: ".debug_aranges", ID: 8, RawID: 9, Addr: 0x0, Size: 0x30},
			{Name: ".rela.debug_aranges", ID: 9, RawID: 10, Addr: 0x0, Size: 0x30},
			{Name: ".debug_line", ID: 10, RawID: 11, Addr: 0x0, Size: 0x45},
			{Name: ".rela.debug_line", ID: 11, RawID: 12, Addr: 0x0, Size: 0x18},
			{Name: ".debug_str", ID: 12, RawID: 13, Addr: 0x0, Size: 0x122},
			{Name: ".comment", ID: 13, RawID: 14, Addr: 0x0, Size: 0x2c},
			{Name: ".note.GNU-stack", ID: 14, RawID: 15, Addr: 0x0, Size: 0x0},
			{Name: ".note.gnu.property", ID: 15, RawID: 16, Addr: 0x0, Size: 0x20},
			{Name: ".eh_frame", ID: 16, RawID: 17, Addr: 0x0, Size: 0x38},
			{Name: ".rela.eh_frame", ID: 17, RawID: 18, Addr: 0x0, Size: 0x18},
			{Name: ".symtab", ID: 18, RawID: 19, Addr: 0x0, Size: 0x1b0},
			{Name: ".strtab", ID: 19, RawID: 20, Addr: 0x0, Size: 0x29},
			{Name: ".shstrtab", ID: 20, RawID: 21, Addr: 0x0, Size: 0xc3},
		},
		nSyms:    17,
		textData: parseHex(`f30f1efa554889e54883ec10897dfc488975f0488d3d00000000e800000000b800000000c9c3`),
	},
	{
		path: "hello-gcc10.3.0-AMD64-rel-gz.o",
		arch: arch.AMD64,
		sections: []Section{
			{Name: ".text", ID: 0, RawID: 1, Addr: 0x0, Size: 0x26},
			{Name: ".rela.text", ID: 1, RawID: 2, Addr: 0x0, Size: 0x30},
			{Name: ".data", ID: 2, RawID: 3, Addr: 0x0, Size: 0x0},
			{Name: ".bss", ID: 3, RawID: 4, Addr: 0x0, Size: 0x0},
			{Name: ".rodata", ID: 4, RawID: 5, Addr: 0x0, Size: 0x6},
			{Name: ".debug_info", ID: 5, RawID: 6, Addr: 0x0, Size: 0xba},
			{Name: ".rela.debug_info", ID: 6, RawID: 7, Addr: 0x0, Size: 0x1b0},
			{Name: ".debug_abbrev", ID: 7, RawID: 8, Addr: 0x0, Size: 0x62},
			{Name: ".debug_aranges", ID: 8, RawID: 9, Addr: 0x0, Size: 0x30},
			{Name: ".rela.debug_aranges", ID: 9, RawID: 10, Addr: 0x0, Size: 0x30},
			{Name: ".debug_line", ID: 10, RawID: 11, Addr: 0x0, Size: 0x45},
			{Name: ".rela.debug_line", ID: 11, RawID: 12, Addr: 0x0, Size: 0x18},
			{Name: ".debug_str", ID: 12, RawID: 13, Addr: 0x0, Size: 0x12b},
			{Name: ".comment", ID: 13, RawID: 14, Addr: 0x0, Size: 0x2c},
			{Name: ".note.GNU-stack", ID: 14, RawID: 15, Addr: 0x0, Size: 0x0},
			{Name: ".note.gnu.property", ID: 15, RawID: 16, Addr: 0x0, Size: 0x20},
			{Name: ".eh_frame", ID: 16, RawID: 17, Addr: 0x0, Size: 0x38},
			{Name: ".rela.eh_frame", ID: 17, RawID: 18, Addr: 0x0, Size: 0x18},
			{Name: ".symtab", ID: 18, RawID: 19, Addr: 0x0, Size: 0x1b0},
			{Name: ".strtab", ID: 19, RawID: 20, Addr: 0x0, Size: 0x29},
			{Name: ".shstrtab", ID: 20, RawID: 21, Addr: 0x0, Size: 0xc3},
		},
		nSyms:         17,
		textData:      parseHex(`f30f1efa554889e54883ec10897dfc488975f0488d3d00000000e800000000b800000000c9c3`),
		debugInfoData: parseHex(`b60000000400000000000801000000000c00000000000000000000000000000000260000000000000000000000020807000000000204070000000002010800000000020207000000000201060000000002020500000000030405696e74000208050000000004086b0000000201060000000005000000000103055700000000000000000000002600000000000000019cb3000000060000000001030e5700000002916c060000000001031bb30000000291600004086500000000`),
	},
}
