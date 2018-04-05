
//line rfc5424/parser.rl:1
package rfc5424

import (
  "fmt"
  "time"
  "github.com/influxdata/go-syslog/chars"
)
 

//line rfc5424/parser.go:13
const rfc5424_start int = 1
const rfc5424_first_final int = 619
const rfc5424_error int = 0

const rfc5424_en_line int = 623
const rfc5424_en_main int = 1


//line rfc5424/parser.rl:12


func Parse(data string) (*SyslogMessage, error) {
    cs, p, pe, eof := 0, 0, len(data), len(data)

    _ = eof

    cr := chars.NewRepo()

    poss := make(map[string]int, 0)

    err := fmt.Errorf("generic error")

    var prival *Prival
    var version *Version
    var timestamp *time.Time
    var hostname string
    var appname string
    var procid string
    var msgid string

    
//line rfc5424/parser.go:45
	{
	cs = rfc5424_start
	}

//line rfc5424/parser.go:50
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 619:
		goto st_case_619
	case 620:
		goto st_case_620
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 621:
		goto st_case_621
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
	case 87:
		goto st_case_87
	case 88:
		goto st_case_88
	case 89:
		goto st_case_89
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 93:
		goto st_case_93
	case 94:
		goto st_case_94
	case 95:
		goto st_case_95
	case 96:
		goto st_case_96
	case 97:
		goto st_case_97
	case 98:
		goto st_case_98
	case 99:
		goto st_case_99
	case 100:
		goto st_case_100
	case 101:
		goto st_case_101
	case 102:
		goto st_case_102
	case 103:
		goto st_case_103
	case 104:
		goto st_case_104
	case 105:
		goto st_case_105
	case 106:
		goto st_case_106
	case 107:
		goto st_case_107
	case 108:
		goto st_case_108
	case 109:
		goto st_case_109
	case 110:
		goto st_case_110
	case 111:
		goto st_case_111
	case 112:
		goto st_case_112
	case 113:
		goto st_case_113
	case 114:
		goto st_case_114
	case 115:
		goto st_case_115
	case 116:
		goto st_case_116
	case 117:
		goto st_case_117
	case 622:
		goto st_case_622
	case 118:
		goto st_case_118
	case 119:
		goto st_case_119
	case 120:
		goto st_case_120
	case 121:
		goto st_case_121
	case 122:
		goto st_case_122
	case 123:
		goto st_case_123
	case 124:
		goto st_case_124
	case 125:
		goto st_case_125
	case 126:
		goto st_case_126
	case 127:
		goto st_case_127
	case 128:
		goto st_case_128
	case 129:
		goto st_case_129
	case 130:
		goto st_case_130
	case 131:
		goto st_case_131
	case 132:
		goto st_case_132
	case 133:
		goto st_case_133
	case 134:
		goto st_case_134
	case 135:
		goto st_case_135
	case 136:
		goto st_case_136
	case 137:
		goto st_case_137
	case 138:
		goto st_case_138
	case 139:
		goto st_case_139
	case 140:
		goto st_case_140
	case 141:
		goto st_case_141
	case 142:
		goto st_case_142
	case 143:
		goto st_case_143
	case 144:
		goto st_case_144
	case 145:
		goto st_case_145
	case 146:
		goto st_case_146
	case 147:
		goto st_case_147
	case 148:
		goto st_case_148
	case 149:
		goto st_case_149
	case 150:
		goto st_case_150
	case 151:
		goto st_case_151
	case 152:
		goto st_case_152
	case 153:
		goto st_case_153
	case 154:
		goto st_case_154
	case 155:
		goto st_case_155
	case 156:
		goto st_case_156
	case 157:
		goto st_case_157
	case 158:
		goto st_case_158
	case 159:
		goto st_case_159
	case 160:
		goto st_case_160
	case 161:
		goto st_case_161
	case 162:
		goto st_case_162
	case 163:
		goto st_case_163
	case 164:
		goto st_case_164
	case 165:
		goto st_case_165
	case 166:
		goto st_case_166
	case 167:
		goto st_case_167
	case 168:
		goto st_case_168
	case 169:
		goto st_case_169
	case 170:
		goto st_case_170
	case 171:
		goto st_case_171
	case 172:
		goto st_case_172
	case 173:
		goto st_case_173
	case 174:
		goto st_case_174
	case 175:
		goto st_case_175
	case 176:
		goto st_case_176
	case 177:
		goto st_case_177
	case 178:
		goto st_case_178
	case 179:
		goto st_case_179
	case 180:
		goto st_case_180
	case 181:
		goto st_case_181
	case 182:
		goto st_case_182
	case 183:
		goto st_case_183
	case 184:
		goto st_case_184
	case 185:
		goto st_case_185
	case 186:
		goto st_case_186
	case 187:
		goto st_case_187
	case 188:
		goto st_case_188
	case 189:
		goto st_case_189
	case 190:
		goto st_case_190
	case 191:
		goto st_case_191
	case 192:
		goto st_case_192
	case 193:
		goto st_case_193
	case 194:
		goto st_case_194
	case 195:
		goto st_case_195
	case 196:
		goto st_case_196
	case 197:
		goto st_case_197
	case 198:
		goto st_case_198
	case 199:
		goto st_case_199
	case 200:
		goto st_case_200
	case 201:
		goto st_case_201
	case 202:
		goto st_case_202
	case 203:
		goto st_case_203
	case 204:
		goto st_case_204
	case 205:
		goto st_case_205
	case 206:
		goto st_case_206
	case 207:
		goto st_case_207
	case 208:
		goto st_case_208
	case 209:
		goto st_case_209
	case 210:
		goto st_case_210
	case 211:
		goto st_case_211
	case 212:
		goto st_case_212
	case 213:
		goto st_case_213
	case 214:
		goto st_case_214
	case 215:
		goto st_case_215
	case 216:
		goto st_case_216
	case 217:
		goto st_case_217
	case 218:
		goto st_case_218
	case 219:
		goto st_case_219
	case 220:
		goto st_case_220
	case 221:
		goto st_case_221
	case 222:
		goto st_case_222
	case 223:
		goto st_case_223
	case 224:
		goto st_case_224
	case 225:
		goto st_case_225
	case 226:
		goto st_case_226
	case 227:
		goto st_case_227
	case 228:
		goto st_case_228
	case 229:
		goto st_case_229
	case 230:
		goto st_case_230
	case 231:
		goto st_case_231
	case 232:
		goto st_case_232
	case 233:
		goto st_case_233
	case 234:
		goto st_case_234
	case 235:
		goto st_case_235
	case 236:
		goto st_case_236
	case 237:
		goto st_case_237
	case 238:
		goto st_case_238
	case 239:
		goto st_case_239
	case 240:
		goto st_case_240
	case 241:
		goto st_case_241
	case 242:
		goto st_case_242
	case 243:
		goto st_case_243
	case 244:
		goto st_case_244
	case 245:
		goto st_case_245
	case 246:
		goto st_case_246
	case 247:
		goto st_case_247
	case 248:
		goto st_case_248
	case 249:
		goto st_case_249
	case 250:
		goto st_case_250
	case 251:
		goto st_case_251
	case 252:
		goto st_case_252
	case 253:
		goto st_case_253
	case 254:
		goto st_case_254
	case 255:
		goto st_case_255
	case 256:
		goto st_case_256
	case 257:
		goto st_case_257
	case 258:
		goto st_case_258
	case 259:
		goto st_case_259
	case 260:
		goto st_case_260
	case 261:
		goto st_case_261
	case 262:
		goto st_case_262
	case 263:
		goto st_case_263
	case 264:
		goto st_case_264
	case 265:
		goto st_case_265
	case 266:
		goto st_case_266
	case 267:
		goto st_case_267
	case 268:
		goto st_case_268
	case 269:
		goto st_case_269
	case 270:
		goto st_case_270
	case 271:
		goto st_case_271
	case 272:
		goto st_case_272
	case 273:
		goto st_case_273
	case 274:
		goto st_case_274
	case 275:
		goto st_case_275
	case 276:
		goto st_case_276
	case 277:
		goto st_case_277
	case 278:
		goto st_case_278
	case 279:
		goto st_case_279
	case 280:
		goto st_case_280
	case 281:
		goto st_case_281
	case 282:
		goto st_case_282
	case 283:
		goto st_case_283
	case 284:
		goto st_case_284
	case 285:
		goto st_case_285
	case 286:
		goto st_case_286
	case 287:
		goto st_case_287
	case 288:
		goto st_case_288
	case 289:
		goto st_case_289
	case 290:
		goto st_case_290
	case 291:
		goto st_case_291
	case 292:
		goto st_case_292
	case 293:
		goto st_case_293
	case 294:
		goto st_case_294
	case 295:
		goto st_case_295
	case 296:
		goto st_case_296
	case 297:
		goto st_case_297
	case 298:
		goto st_case_298
	case 299:
		goto st_case_299
	case 300:
		goto st_case_300
	case 301:
		goto st_case_301
	case 302:
		goto st_case_302
	case 303:
		goto st_case_303
	case 304:
		goto st_case_304
	case 305:
		goto st_case_305
	case 306:
		goto st_case_306
	case 307:
		goto st_case_307
	case 308:
		goto st_case_308
	case 309:
		goto st_case_309
	case 310:
		goto st_case_310
	case 311:
		goto st_case_311
	case 312:
		goto st_case_312
	case 313:
		goto st_case_313
	case 314:
		goto st_case_314
	case 315:
		goto st_case_315
	case 316:
		goto st_case_316
	case 317:
		goto st_case_317
	case 318:
		goto st_case_318
	case 319:
		goto st_case_319
	case 320:
		goto st_case_320
	case 321:
		goto st_case_321
	case 322:
		goto st_case_322
	case 323:
		goto st_case_323
	case 324:
		goto st_case_324
	case 325:
		goto st_case_325
	case 326:
		goto st_case_326
	case 327:
		goto st_case_327
	case 328:
		goto st_case_328
	case 329:
		goto st_case_329
	case 330:
		goto st_case_330
	case 331:
		goto st_case_331
	case 332:
		goto st_case_332
	case 333:
		goto st_case_333
	case 334:
		goto st_case_334
	case 335:
		goto st_case_335
	case 336:
		goto st_case_336
	case 337:
		goto st_case_337
	case 338:
		goto st_case_338
	case 339:
		goto st_case_339
	case 340:
		goto st_case_340
	case 341:
		goto st_case_341
	case 342:
		goto st_case_342
	case 343:
		goto st_case_343
	case 344:
		goto st_case_344
	case 345:
		goto st_case_345
	case 346:
		goto st_case_346
	case 347:
		goto st_case_347
	case 348:
		goto st_case_348
	case 349:
		goto st_case_349
	case 350:
		goto st_case_350
	case 351:
		goto st_case_351
	case 352:
		goto st_case_352
	case 353:
		goto st_case_353
	case 354:
		goto st_case_354
	case 355:
		goto st_case_355
	case 356:
		goto st_case_356
	case 357:
		goto st_case_357
	case 358:
		goto st_case_358
	case 359:
		goto st_case_359
	case 360:
		goto st_case_360
	case 361:
		goto st_case_361
	case 362:
		goto st_case_362
	case 363:
		goto st_case_363
	case 364:
		goto st_case_364
	case 365:
		goto st_case_365
	case 366:
		goto st_case_366
	case 367:
		goto st_case_367
	case 368:
		goto st_case_368
	case 369:
		goto st_case_369
	case 370:
		goto st_case_370
	case 371:
		goto st_case_371
	case 372:
		goto st_case_372
	case 373:
		goto st_case_373
	case 374:
		goto st_case_374
	case 375:
		goto st_case_375
	case 376:
		goto st_case_376
	case 377:
		goto st_case_377
	case 378:
		goto st_case_378
	case 379:
		goto st_case_379
	case 380:
		goto st_case_380
	case 381:
		goto st_case_381
	case 382:
		goto st_case_382
	case 383:
		goto st_case_383
	case 384:
		goto st_case_384
	case 385:
		goto st_case_385
	case 386:
		goto st_case_386
	case 387:
		goto st_case_387
	case 388:
		goto st_case_388
	case 389:
		goto st_case_389
	case 390:
		goto st_case_390
	case 391:
		goto st_case_391
	case 392:
		goto st_case_392
	case 393:
		goto st_case_393
	case 394:
		goto st_case_394
	case 395:
		goto st_case_395
	case 396:
		goto st_case_396
	case 397:
		goto st_case_397
	case 398:
		goto st_case_398
	case 399:
		goto st_case_399
	case 400:
		goto st_case_400
	case 401:
		goto st_case_401
	case 402:
		goto st_case_402
	case 403:
		goto st_case_403
	case 404:
		goto st_case_404
	case 405:
		goto st_case_405
	case 406:
		goto st_case_406
	case 407:
		goto st_case_407
	case 408:
		goto st_case_408
	case 409:
		goto st_case_409
	case 410:
		goto st_case_410
	case 411:
		goto st_case_411
	case 412:
		goto st_case_412
	case 413:
		goto st_case_413
	case 414:
		goto st_case_414
	case 415:
		goto st_case_415
	case 416:
		goto st_case_416
	case 417:
		goto st_case_417
	case 418:
		goto st_case_418
	case 419:
		goto st_case_419
	case 420:
		goto st_case_420
	case 421:
		goto st_case_421
	case 422:
		goto st_case_422
	case 423:
		goto st_case_423
	case 424:
		goto st_case_424
	case 425:
		goto st_case_425
	case 426:
		goto st_case_426
	case 427:
		goto st_case_427
	case 428:
		goto st_case_428
	case 429:
		goto st_case_429
	case 430:
		goto st_case_430
	case 431:
		goto st_case_431
	case 432:
		goto st_case_432
	case 433:
		goto st_case_433
	case 434:
		goto st_case_434
	case 435:
		goto st_case_435
	case 436:
		goto st_case_436
	case 437:
		goto st_case_437
	case 438:
		goto st_case_438
	case 439:
		goto st_case_439
	case 440:
		goto st_case_440
	case 441:
		goto st_case_441
	case 442:
		goto st_case_442
	case 443:
		goto st_case_443
	case 444:
		goto st_case_444
	case 445:
		goto st_case_445
	case 446:
		goto st_case_446
	case 447:
		goto st_case_447
	case 448:
		goto st_case_448
	case 449:
		goto st_case_449
	case 450:
		goto st_case_450
	case 451:
		goto st_case_451
	case 452:
		goto st_case_452
	case 453:
		goto st_case_453
	case 454:
		goto st_case_454
	case 455:
		goto st_case_455
	case 456:
		goto st_case_456
	case 457:
		goto st_case_457
	case 458:
		goto st_case_458
	case 459:
		goto st_case_459
	case 460:
		goto st_case_460
	case 461:
		goto st_case_461
	case 462:
		goto st_case_462
	case 463:
		goto st_case_463
	case 464:
		goto st_case_464
	case 465:
		goto st_case_465
	case 466:
		goto st_case_466
	case 467:
		goto st_case_467
	case 468:
		goto st_case_468
	case 469:
		goto st_case_469
	case 470:
		goto st_case_470
	case 471:
		goto st_case_471
	case 472:
		goto st_case_472
	case 473:
		goto st_case_473
	case 474:
		goto st_case_474
	case 475:
		goto st_case_475
	case 476:
		goto st_case_476
	case 477:
		goto st_case_477
	case 478:
		goto st_case_478
	case 479:
		goto st_case_479
	case 480:
		goto st_case_480
	case 481:
		goto st_case_481
	case 482:
		goto st_case_482
	case 483:
		goto st_case_483
	case 484:
		goto st_case_484
	case 485:
		goto st_case_485
	case 486:
		goto st_case_486
	case 487:
		goto st_case_487
	case 488:
		goto st_case_488
	case 489:
		goto st_case_489
	case 490:
		goto st_case_490
	case 491:
		goto st_case_491
	case 492:
		goto st_case_492
	case 493:
		goto st_case_493
	case 494:
		goto st_case_494
	case 495:
		goto st_case_495
	case 496:
		goto st_case_496
	case 497:
		goto st_case_497
	case 498:
		goto st_case_498
	case 499:
		goto st_case_499
	case 500:
		goto st_case_500
	case 501:
		goto st_case_501
	case 502:
		goto st_case_502
	case 503:
		goto st_case_503
	case 504:
		goto st_case_504
	case 505:
		goto st_case_505
	case 506:
		goto st_case_506
	case 507:
		goto st_case_507
	case 508:
		goto st_case_508
	case 509:
		goto st_case_509
	case 510:
		goto st_case_510
	case 511:
		goto st_case_511
	case 512:
		goto st_case_512
	case 513:
		goto st_case_513
	case 514:
		goto st_case_514
	case 515:
		goto st_case_515
	case 516:
		goto st_case_516
	case 517:
		goto st_case_517
	case 518:
		goto st_case_518
	case 519:
		goto st_case_519
	case 520:
		goto st_case_520
	case 521:
		goto st_case_521
	case 522:
		goto st_case_522
	case 523:
		goto st_case_523
	case 524:
		goto st_case_524
	case 525:
		goto st_case_525
	case 526:
		goto st_case_526
	case 527:
		goto st_case_527
	case 528:
		goto st_case_528
	case 529:
		goto st_case_529
	case 530:
		goto st_case_530
	case 531:
		goto st_case_531
	case 532:
		goto st_case_532
	case 533:
		goto st_case_533
	case 534:
		goto st_case_534
	case 535:
		goto st_case_535
	case 536:
		goto st_case_536
	case 537:
		goto st_case_537
	case 538:
		goto st_case_538
	case 539:
		goto st_case_539
	case 540:
		goto st_case_540
	case 541:
		goto st_case_541
	case 542:
		goto st_case_542
	case 543:
		goto st_case_543
	case 544:
		goto st_case_544
	case 545:
		goto st_case_545
	case 546:
		goto st_case_546
	case 547:
		goto st_case_547
	case 548:
		goto st_case_548
	case 549:
		goto st_case_549
	case 550:
		goto st_case_550
	case 551:
		goto st_case_551
	case 552:
		goto st_case_552
	case 553:
		goto st_case_553
	case 554:
		goto st_case_554
	case 555:
		goto st_case_555
	case 556:
		goto st_case_556
	case 557:
		goto st_case_557
	case 558:
		goto st_case_558
	case 559:
		goto st_case_559
	case 560:
		goto st_case_560
	case 561:
		goto st_case_561
	case 562:
		goto st_case_562
	case 563:
		goto st_case_563
	case 564:
		goto st_case_564
	case 565:
		goto st_case_565
	case 566:
		goto st_case_566
	case 567:
		goto st_case_567
	case 568:
		goto st_case_568
	case 569:
		goto st_case_569
	case 570:
		goto st_case_570
	case 571:
		goto st_case_571
	case 572:
		goto st_case_572
	case 573:
		goto st_case_573
	case 574:
		goto st_case_574
	case 575:
		goto st_case_575
	case 576:
		goto st_case_576
	case 577:
		goto st_case_577
	case 578:
		goto st_case_578
	case 579:
		goto st_case_579
	case 580:
		goto st_case_580
	case 581:
		goto st_case_581
	case 582:
		goto st_case_582
	case 583:
		goto st_case_583
	case 584:
		goto st_case_584
	case 585:
		goto st_case_585
	case 586:
		goto st_case_586
	case 587:
		goto st_case_587
	case 588:
		goto st_case_588
	case 589:
		goto st_case_589
	case 590:
		goto st_case_590
	case 591:
		goto st_case_591
	case 592:
		goto st_case_592
	case 593:
		goto st_case_593
	case 594:
		goto st_case_594
	case 595:
		goto st_case_595
	case 596:
		goto st_case_596
	case 597:
		goto st_case_597
	case 598:
		goto st_case_598
	case 599:
		goto st_case_599
	case 600:
		goto st_case_600
	case 601:
		goto st_case_601
	case 602:
		goto st_case_602
	case 603:
		goto st_case_603
	case 604:
		goto st_case_604
	case 605:
		goto st_case_605
	case 606:
		goto st_case_606
	case 607:
		goto st_case_607
	case 608:
		goto st_case_608
	case 609:
		goto st_case_609
	case 610:
		goto st_case_610
	case 611:
		goto st_case_611
	case 612:
		goto st_case_612
	case 613:
		goto st_case_613
	case 614:
		goto st_case_614
	case 615:
		goto st_case_615
	case 616:
		goto st_case_616
	case 617:
		goto st_case_617
	case 618:
		goto st_case_618
	case 623:
		goto st_case_623
	}
	goto st_out
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		if data[p] == 60 {
			goto st2
		}
		goto st0
tr9:
//line rfc5424/machine.rl:49

    err = fmt.Errorf("error parsing <nilvalue>");

	goto st0
//line rfc5424/parser.go:1321
st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 48:
			goto tr2
		case 49:
			goto tr3
		}
		if 50 <= data[p] && data[p] <= 57 {
			goto tr4
		}
		goto st0
tr2:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line rfc5424/parser.go:1352
		if data[p] == 62 {
			goto tr5
		}
		goto st0
tr5:
//line rfc5424/machine.rl:11

    prival = NewPrival(*cr.ReduceToInt(chars.UTF8DecimalCodePointsToInt))

	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line rfc5424/parser.go:1368
		if 49 <= data[p] && data[p] <= 57 {
			goto tr6
		}
		goto st0
tr6:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line rfc5424/parser.go:1384
		if data[p] == 32 {
			goto tr7
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr8
		}
		goto st0
tr7:
//line rfc5424/machine.rl:15

    version = NewVersion(*cr.ReduceToInt(chars.UTF8DecimalCodePointsToInt))

	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line rfc5424/parser.go:1403
		if data[p] == 45 {
			goto st7
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr11
		}
		goto tr9
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if data[p] == 32 {
			goto st8
		}
		goto st0
tr617:
//line rfc5424/machine.rl:40

    if t, e := time.Parse(time.RFC3339Nano, data[poss["timestamp:ini"]:p]); e != nil {
        err = fmt.Errorf("error %s [col %d:%d]", e, poss["timestamp:ini"], p);
        p--
 {goto st623 }
    } else {
        timestamp = &t
    }

	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line rfc5424/parser.go:1437
		if data[p] == 45 {
			goto tr14
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr13
		}
		goto tr9
tr14:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st9
tr13:
//line rfc5424/machine.rl:49

    err = fmt.Errorf("error parsing <nilvalue>");

//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line rfc5424/parser.go:1466
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr16
		}
		goto st0
tr15:
//line rfc5424/machine.rl:19

    hostname = *cr.ReduceToString(chars.UTF8DecimalCodePointsToString)

	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line rfc5424/parser.go:1485
		if data[p] == 45 {
			goto tr18
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr17
		}
		goto tr9
tr18:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st11
tr17:
//line rfc5424/machine.rl:49

    err = fmt.Errorf("error parsing <nilvalue>");

//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line rfc5424/parser.go:1514
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr20
		}
		goto st0
tr19:
//line rfc5424/machine.rl:23

    appname = *cr.ReduceToString(chars.UTF8DecimalCodePointsToString)

	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line rfc5424/parser.go:1533
		if data[p] == 45 {
			goto tr22
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr21
		}
		goto tr9
tr22:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st13
tr21:
//line rfc5424/machine.rl:49

    err = fmt.Errorf("error parsing <nilvalue>");

//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line rfc5424/parser.go:1562
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr24
		}
		goto st0
tr23:
//line rfc5424/machine.rl:27

    fmt.Println("set")
    //procid = *cr.ReduceToString(chars.UTF8DecimalCodePointsToString)

	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line rfc5424/parser.go:1582
		if data[p] == 45 {
			goto tr26
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr25
		}
		goto tr9
tr26:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st15
tr25:
//line rfc5424/machine.rl:49

    err = fmt.Errorf("error parsing <nilvalue>");

//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line rfc5424/parser.go:1611
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr28
		}
		goto st0
tr27:
//line rfc5424/machine.rl:32

    msgid = *cr.ReduceToString(chars.UTF8DecimalCodePointsToString)

	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line rfc5424/parser.go:1630
		switch data[p] {
		case 45:
			goto st619
		case 91:
			goto tr30
		}
		goto tr9
	st619:
		if p++; p == pe {
			goto _test_eof619
		}
	st_case_619:
		if data[p] == 32 {
			goto st620
		}
		goto st0
	st620:
		if p++; p == pe {
			goto _test_eof620
		}
	st_case_620:
		goto st620
tr30:
//line rfc5424/machine.rl:49

    err = fmt.Errorf("error parsing <nilvalue>");

	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line rfc5424/parser.go:1664
		if data[p] == 33 {
			goto st18
		}
		switch {
		case data[p] < 62:
			if 35 <= data[p] && data[p] <= 60 {
				goto st18
			}
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st18
			}
		default:
			goto st18
		}
		goto st0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st87
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st87
			}
		case data[p] >= 35:
			goto st87
		}
		goto st0
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if data[p] == 33 {
			goto st20
		}
		switch {
		case data[p] < 62:
			if 35 <= data[p] && data[p] <= 60 {
				goto st20
			}
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st20
			}
		default:
			goto st20
		}
		goto st0
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		switch data[p] {
		case 33:
			goto st21
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st21
			}
		case data[p] >= 35:
			goto st21
		}
		goto st0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		switch data[p] {
		case 33:
			goto st22
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st22
			}
		case data[p] >= 35:
			goto st22
		}
		goto st0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch data[p] {
		case 33:
			goto st23
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st23
			}
		case data[p] >= 35:
			goto st23
		}
		goto st0
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		switch data[p] {
		case 33:
			goto st24
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st24
			}
		case data[p] >= 35:
			goto st24
		}
		goto st0
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		switch data[p] {
		case 33:
			goto st25
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st25
			}
		case data[p] >= 35:
			goto st25
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch data[p] {
		case 33:
			goto st26
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st26
			}
		case data[p] >= 35:
			goto st26
		}
		goto st0
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		switch data[p] {
		case 33:
			goto st27
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st27
			}
		case data[p] >= 35:
			goto st27
		}
		goto st0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch data[p] {
		case 33:
			goto st28
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st28
			}
		case data[p] >= 35:
			goto st28
		}
		goto st0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch data[p] {
		case 33:
			goto st29
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st29
			}
		case data[p] >= 35:
			goto st29
		}
		goto st0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch data[p] {
		case 33:
			goto st30
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st30
			}
		case data[p] >= 35:
			goto st30
		}
		goto st0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 33:
			goto st31
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st31
			}
		case data[p] >= 35:
			goto st31
		}
		goto st0
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch data[p] {
		case 33:
			goto st32
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st32
			}
		case data[p] >= 35:
			goto st32
		}
		goto st0
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch data[p] {
		case 33:
			goto st33
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st33
			}
		case data[p] >= 35:
			goto st33
		}
		goto st0
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 33:
			goto st34
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st34
			}
		case data[p] >= 35:
			goto st34
		}
		goto st0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch data[p] {
		case 33:
			goto st35
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st35
			}
		case data[p] >= 35:
			goto st35
		}
		goto st0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch data[p] {
		case 33:
			goto st36
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st36
			}
		case data[p] >= 35:
			goto st36
		}
		goto st0
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		switch data[p] {
		case 33:
			goto st37
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st37
			}
		case data[p] >= 35:
			goto st37
		}
		goto st0
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		switch data[p] {
		case 33:
			goto st38
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st38
			}
		case data[p] >= 35:
			goto st38
		}
		goto st0
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		switch data[p] {
		case 33:
			goto st39
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st39
			}
		case data[p] >= 35:
			goto st39
		}
		goto st0
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		switch data[p] {
		case 33:
			goto st40
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st40
			}
		case data[p] >= 35:
			goto st40
		}
		goto st0
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		switch data[p] {
		case 33:
			goto st41
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st41
			}
		case data[p] >= 35:
			goto st41
		}
		goto st0
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		switch data[p] {
		case 33:
			goto st42
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st42
			}
		case data[p] >= 35:
			goto st42
		}
		goto st0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		switch data[p] {
		case 33:
			goto st43
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st43
			}
		case data[p] >= 35:
			goto st43
		}
		goto st0
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		switch data[p] {
		case 33:
			goto st44
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st44
			}
		case data[p] >= 35:
			goto st44
		}
		goto st0
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		switch data[p] {
		case 33:
			goto st45
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st45
			}
		case data[p] >= 35:
			goto st45
		}
		goto st0
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		switch data[p] {
		case 33:
			goto st46
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st46
			}
		case data[p] >= 35:
			goto st46
		}
		goto st0
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		switch data[p] {
		case 33:
			goto st47
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st47
			}
		case data[p] >= 35:
			goto st47
		}
		goto st0
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		switch data[p] {
		case 33:
			goto st48
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st48
			}
		case data[p] >= 35:
			goto st48
		}
		goto st0
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		switch data[p] {
		case 33:
			goto st49
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st49
			}
		case data[p] >= 35:
			goto st49
		}
		goto st0
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		switch data[p] {
		case 33:
			goto st50
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st50
			}
		case data[p] >= 35:
			goto st50
		}
		goto st0
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		switch data[p] {
		case 33:
			goto st51
		case 61:
			goto st52
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st51
			}
		case data[p] >= 35:
			goto st51
		}
		goto st0
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		if data[p] == 61 {
			goto st52
		}
		goto st0
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		if data[p] == 34 {
			goto st53
		}
		goto st0
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		if data[p] == 34 {
			goto st54
		}
		goto st53
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		goto st53
	st621:
		if p++; p == pe {
			goto _test_eof621
		}
	st_case_621:
		switch data[p] {
		case 32:
			goto st620
		case 34:
			goto st54
		case 91:
			goto st55
		}
		goto st53
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		if data[p] == 34 {
			goto st54
		}
		switch {
		case data[p] < 62:
			if 33 <= data[p] && data[p] <= 60 {
				goto st56
			}
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st56
			}
		default:
			goto st56
		}
		goto st53
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st57
			}
		case data[p] >= 33:
			goto st57
		}
		goto st53
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st58
			}
		case data[p] >= 33:
			goto st58
		}
		goto st53
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st59
			}
		case data[p] >= 33:
			goto st59
		}
		goto st53
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st60
			}
		case data[p] >= 33:
			goto st60
		}
		goto st53
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st61
			}
		case data[p] >= 33:
			goto st61
		}
		goto st53
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st62
			}
		case data[p] >= 33:
			goto st62
		}
		goto st53
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st63
			}
		case data[p] >= 33:
			goto st63
		}
		goto st53
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st64
			}
		case data[p] >= 33:
			goto st64
		}
		goto st53
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st65
			}
		case data[p] >= 33:
			goto st65
		}
		goto st53
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st66
			}
		case data[p] >= 33:
			goto st66
		}
		goto st53
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st67
			}
		case data[p] >= 33:
			goto st67
		}
		goto st53
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st68
			}
		case data[p] >= 33:
			goto st68
		}
		goto st53
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st69
			}
		case data[p] >= 33:
			goto st69
		}
		goto st53
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st70
			}
		case data[p] >= 33:
			goto st70
		}
		goto st53
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st71
			}
		case data[p] >= 33:
			goto st71
		}
		goto st53
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st72
			}
		case data[p] >= 33:
			goto st72
		}
		goto st53
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st73
			}
		case data[p] >= 33:
			goto st73
		}
		goto st53
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st74
			}
		case data[p] >= 33:
			goto st74
		}
		goto st53
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st75
			}
		case data[p] >= 33:
			goto st75
		}
		goto st53
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st76
			}
		case data[p] >= 33:
			goto st76
		}
		goto st53
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st77
			}
		case data[p] >= 33:
			goto st77
		}
		goto st53
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st78
			}
		case data[p] >= 33:
			goto st78
		}
		goto st53
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st79
			}
		case data[p] >= 33:
			goto st79
		}
		goto st53
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st80
			}
		case data[p] >= 33:
			goto st80
		}
		goto st53
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st81
			}
		case data[p] >= 33:
			goto st81
		}
		goto st53
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st82
			}
		case data[p] >= 33:
			goto st82
		}
		goto st53
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st83
			}
		case data[p] >= 33:
			goto st83
		}
		goto st53
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st84
			}
		case data[p] >= 33:
			goto st84
		}
		goto st53
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st85
			}
		case data[p] >= 33:
			goto st85
		}
		goto st53
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		switch data[p] {
		case 34:
			goto st54
		case 93:
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st86
			}
		case data[p] >= 33:
			goto st86
		}
		goto st53
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		if data[p] == 93 {
			goto st621
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st54
			}
		case data[p] >= 33:
			goto st54
		}
		goto st53
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st88
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st88
			}
		case data[p] >= 35:
			goto st88
		}
		goto st0
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st89
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st89
			}
		case data[p] >= 35:
			goto st89
		}
		goto st0
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st90
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st90
			}
		case data[p] >= 35:
			goto st90
		}
		goto st0
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st91
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st91
			}
		case data[p] >= 35:
			goto st91
		}
		goto st0
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st92
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st92
			}
		case data[p] >= 35:
			goto st92
		}
		goto st0
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st93
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st93
			}
		case data[p] >= 35:
			goto st93
		}
		goto st0
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st94
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st94
			}
		case data[p] >= 35:
			goto st94
		}
		goto st0
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st95
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st95
			}
		case data[p] >= 35:
			goto st95
		}
		goto st0
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st96
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st96
			}
		case data[p] >= 35:
			goto st96
		}
		goto st0
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st97
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st97
			}
		case data[p] >= 35:
			goto st97
		}
		goto st0
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st98
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st98
			}
		case data[p] >= 35:
			goto st98
		}
		goto st0
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st99
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st99
			}
		case data[p] >= 35:
			goto st99
		}
		goto st0
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st100
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st100
			}
		case data[p] >= 35:
			goto st100
		}
		goto st0
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st101
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st101
			}
		case data[p] >= 35:
			goto st101
		}
		goto st0
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st102
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st102
			}
		case data[p] >= 35:
			goto st102
		}
		goto st0
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st103
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st103
			}
		case data[p] >= 35:
			goto st103
		}
		goto st0
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st104
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st104
			}
		case data[p] >= 35:
			goto st104
		}
		goto st0
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st105
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st105
			}
		case data[p] >= 35:
			goto st105
		}
		goto st0
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st106
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st106
			}
		case data[p] >= 35:
			goto st106
		}
		goto st0
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st107
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st107
			}
		case data[p] >= 35:
			goto st107
		}
		goto st0
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st108
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st108
			}
		case data[p] >= 35:
			goto st108
		}
		goto st0
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st109
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st109
			}
		case data[p] >= 35:
			goto st109
		}
		goto st0
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st110
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st110
			}
		case data[p] >= 35:
			goto st110
		}
		goto st0
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st111
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st111
			}
		case data[p] >= 35:
			goto st111
		}
		goto st0
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st112
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st112
			}
		case data[p] >= 35:
			goto st112
		}
		goto st0
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st113
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st113
			}
		case data[p] >= 35:
			goto st113
		}
		goto st0
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st114
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st114
			}
		case data[p] >= 35:
			goto st114
		}
		goto st0
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st115
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st115
			}
		case data[p] >= 35:
			goto st115
		}
		goto st0
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st116
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st116
			}
		case data[p] >= 35:
			goto st116
		}
		goto st0
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		switch data[p] {
		case 32:
			goto st19
		case 33:
			goto st117
		case 93:
			goto st622
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st117
			}
		case data[p] >= 35:
			goto st117
		}
		goto st0
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		switch data[p] {
		case 32:
			goto st19
		case 93:
			goto st622
		}
		goto st0
	st622:
		if p++; p == pe {
			goto _test_eof622
		}
	st_case_622:
		switch data[p] {
		case 32:
			goto st620
		case 91:
			goto st17
		}
		goto st0
tr28:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st118
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
//line rfc5424/parser.go:3730
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr132
		}
		goto st0
tr132:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st119
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
//line rfc5424/parser.go:3749
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr133
		}
		goto st0
tr133:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st120
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
//line rfc5424/parser.go:3768
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr134
		}
		goto st0
tr134:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st121
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
//line rfc5424/parser.go:3787
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr135
		}
		goto st0
tr135:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st122
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
//line rfc5424/parser.go:3806
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr136
		}
		goto st0
tr136:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st123
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
//line rfc5424/parser.go:3825
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr137
		}
		goto st0
tr137:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st124
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
//line rfc5424/parser.go:3844
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr138
		}
		goto st0
tr138:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st125
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
//line rfc5424/parser.go:3863
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr139
		}
		goto st0
tr139:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st126
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
//line rfc5424/parser.go:3882
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr140
		}
		goto st0
tr140:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st127
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
//line rfc5424/parser.go:3901
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr141
		}
		goto st0
tr141:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st128
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
//line rfc5424/parser.go:3920
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr142
		}
		goto st0
tr142:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st129
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
//line rfc5424/parser.go:3939
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr143
		}
		goto st0
tr143:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st130
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
//line rfc5424/parser.go:3958
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr144
		}
		goto st0
tr144:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st131
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
//line rfc5424/parser.go:3977
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr145
		}
		goto st0
tr145:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st132
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
//line rfc5424/parser.go:3996
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr146
		}
		goto st0
tr146:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st133
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
//line rfc5424/parser.go:4015
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr147
		}
		goto st0
tr147:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st134
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
//line rfc5424/parser.go:4034
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr148
		}
		goto st0
tr148:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st135
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
//line rfc5424/parser.go:4053
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr149
		}
		goto st0
tr149:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st136
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
//line rfc5424/parser.go:4072
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr150
		}
		goto st0
tr150:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st137
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
//line rfc5424/parser.go:4091
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr151
		}
		goto st0
tr151:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st138
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
//line rfc5424/parser.go:4110
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr152
		}
		goto st0
tr152:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st139
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
//line rfc5424/parser.go:4129
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr153
		}
		goto st0
tr153:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st140
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
//line rfc5424/parser.go:4148
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr154
		}
		goto st0
tr154:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st141
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
//line rfc5424/parser.go:4167
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr155
		}
		goto st0
tr155:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st142
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
//line rfc5424/parser.go:4186
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr156
		}
		goto st0
tr156:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st143
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
//line rfc5424/parser.go:4205
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr157
		}
		goto st0
tr157:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st144
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
//line rfc5424/parser.go:4224
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr158
		}
		goto st0
tr158:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st145
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
//line rfc5424/parser.go:4243
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr159
		}
		goto st0
tr159:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st146
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
//line rfc5424/parser.go:4262
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr160
		}
		goto st0
tr160:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st147
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
//line rfc5424/parser.go:4281
		if data[p] == 32 {
			goto tr27
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr161
		}
		goto st0
tr161:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st148
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
//line rfc5424/parser.go:4300
		if data[p] == 32 {
			goto tr27
		}
		goto st0
tr24:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st149
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
//line rfc5424/parser.go:4316
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr162
		}
		goto st0
tr162:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st150
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
//line rfc5424/parser.go:4335
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr163
		}
		goto st0
tr163:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st151
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
//line rfc5424/parser.go:4354
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr164
		}
		goto st0
tr164:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st152
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
//line rfc5424/parser.go:4373
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr165
		}
		goto st0
tr165:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st153
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
//line rfc5424/parser.go:4392
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr166
		}
		goto st0
tr166:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st154
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
//line rfc5424/parser.go:4411
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr167
		}
		goto st0
tr167:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st155
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
//line rfc5424/parser.go:4430
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr168
		}
		goto st0
tr168:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st156
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
//line rfc5424/parser.go:4449
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr169
		}
		goto st0
tr169:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st157
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
//line rfc5424/parser.go:4468
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr170
		}
		goto st0
tr170:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st158
	st158:
		if p++; p == pe {
			goto _test_eof158
		}
	st_case_158:
//line rfc5424/parser.go:4487
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr171
		}
		goto st0
tr171:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st159
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
//line rfc5424/parser.go:4506
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr172
		}
		goto st0
tr172:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st160
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
//line rfc5424/parser.go:4525
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr173
		}
		goto st0
tr173:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st161
	st161:
		if p++; p == pe {
			goto _test_eof161
		}
	st_case_161:
//line rfc5424/parser.go:4544
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr174
		}
		goto st0
tr174:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st162
	st162:
		if p++; p == pe {
			goto _test_eof162
		}
	st_case_162:
//line rfc5424/parser.go:4563
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr175
		}
		goto st0
tr175:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st163
	st163:
		if p++; p == pe {
			goto _test_eof163
		}
	st_case_163:
//line rfc5424/parser.go:4582
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr176
		}
		goto st0
tr176:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st164
	st164:
		if p++; p == pe {
			goto _test_eof164
		}
	st_case_164:
//line rfc5424/parser.go:4601
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr177
		}
		goto st0
tr177:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st165
	st165:
		if p++; p == pe {
			goto _test_eof165
		}
	st_case_165:
//line rfc5424/parser.go:4620
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr178
		}
		goto st0
tr178:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st166
	st166:
		if p++; p == pe {
			goto _test_eof166
		}
	st_case_166:
//line rfc5424/parser.go:4639
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr179
		}
		goto st0
tr179:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st167
	st167:
		if p++; p == pe {
			goto _test_eof167
		}
	st_case_167:
//line rfc5424/parser.go:4658
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr180
		}
		goto st0
tr180:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st168
	st168:
		if p++; p == pe {
			goto _test_eof168
		}
	st_case_168:
//line rfc5424/parser.go:4677
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr181
		}
		goto st0
tr181:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st169
	st169:
		if p++; p == pe {
			goto _test_eof169
		}
	st_case_169:
//line rfc5424/parser.go:4696
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr182
		}
		goto st0
tr182:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st170
	st170:
		if p++; p == pe {
			goto _test_eof170
		}
	st_case_170:
//line rfc5424/parser.go:4715
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr183
		}
		goto st0
tr183:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st171
	st171:
		if p++; p == pe {
			goto _test_eof171
		}
	st_case_171:
//line rfc5424/parser.go:4734
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr184
		}
		goto st0
tr184:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st172
	st172:
		if p++; p == pe {
			goto _test_eof172
		}
	st_case_172:
//line rfc5424/parser.go:4753
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr185
		}
		goto st0
tr185:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st173
	st173:
		if p++; p == pe {
			goto _test_eof173
		}
	st_case_173:
//line rfc5424/parser.go:4772
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr186
		}
		goto st0
tr186:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st174
	st174:
		if p++; p == pe {
			goto _test_eof174
		}
	st_case_174:
//line rfc5424/parser.go:4791
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr187
		}
		goto st0
tr187:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st175
	st175:
		if p++; p == pe {
			goto _test_eof175
		}
	st_case_175:
//line rfc5424/parser.go:4810
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr188
		}
		goto st0
tr188:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st176
	st176:
		if p++; p == pe {
			goto _test_eof176
		}
	st_case_176:
//line rfc5424/parser.go:4829
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr189
		}
		goto st0
tr189:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st177
	st177:
		if p++; p == pe {
			goto _test_eof177
		}
	st_case_177:
//line rfc5424/parser.go:4848
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr190
		}
		goto st0
tr190:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st178
	st178:
		if p++; p == pe {
			goto _test_eof178
		}
	st_case_178:
//line rfc5424/parser.go:4867
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr191
		}
		goto st0
tr191:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st179
	st179:
		if p++; p == pe {
			goto _test_eof179
		}
	st_case_179:
//line rfc5424/parser.go:4886
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr192
		}
		goto st0
tr192:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st180
	st180:
		if p++; p == pe {
			goto _test_eof180
		}
	st_case_180:
//line rfc5424/parser.go:4905
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr193
		}
		goto st0
tr193:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st181
	st181:
		if p++; p == pe {
			goto _test_eof181
		}
	st_case_181:
//line rfc5424/parser.go:4924
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr194
		}
		goto st0
tr194:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st182
	st182:
		if p++; p == pe {
			goto _test_eof182
		}
	st_case_182:
//line rfc5424/parser.go:4943
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr195
		}
		goto st0
tr195:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st183
	st183:
		if p++; p == pe {
			goto _test_eof183
		}
	st_case_183:
//line rfc5424/parser.go:4962
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr196
		}
		goto st0
tr196:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st184
	st184:
		if p++; p == pe {
			goto _test_eof184
		}
	st_case_184:
//line rfc5424/parser.go:4981
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr197
		}
		goto st0
tr197:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st185
	st185:
		if p++; p == pe {
			goto _test_eof185
		}
	st_case_185:
//line rfc5424/parser.go:5000
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr198
		}
		goto st0
tr198:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st186
	st186:
		if p++; p == pe {
			goto _test_eof186
		}
	st_case_186:
//line rfc5424/parser.go:5019
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr199
		}
		goto st0
tr199:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st187
	st187:
		if p++; p == pe {
			goto _test_eof187
		}
	st_case_187:
//line rfc5424/parser.go:5038
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr200
		}
		goto st0
tr200:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st188
	st188:
		if p++; p == pe {
			goto _test_eof188
		}
	st_case_188:
//line rfc5424/parser.go:5057
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr201
		}
		goto st0
tr201:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st189
	st189:
		if p++; p == pe {
			goto _test_eof189
		}
	st_case_189:
//line rfc5424/parser.go:5076
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr202
		}
		goto st0
tr202:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st190
	st190:
		if p++; p == pe {
			goto _test_eof190
		}
	st_case_190:
//line rfc5424/parser.go:5095
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr203
		}
		goto st0
tr203:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st191
	st191:
		if p++; p == pe {
			goto _test_eof191
		}
	st_case_191:
//line rfc5424/parser.go:5114
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr204
		}
		goto st0
tr204:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st192
	st192:
		if p++; p == pe {
			goto _test_eof192
		}
	st_case_192:
//line rfc5424/parser.go:5133
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr205
		}
		goto st0
tr205:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st193
	st193:
		if p++; p == pe {
			goto _test_eof193
		}
	st_case_193:
//line rfc5424/parser.go:5152
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr206
		}
		goto st0
tr206:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st194
	st194:
		if p++; p == pe {
			goto _test_eof194
		}
	st_case_194:
//line rfc5424/parser.go:5171
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr207
		}
		goto st0
tr207:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st195
	st195:
		if p++; p == pe {
			goto _test_eof195
		}
	st_case_195:
//line rfc5424/parser.go:5190
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr208
		}
		goto st0
tr208:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st196
	st196:
		if p++; p == pe {
			goto _test_eof196
		}
	st_case_196:
//line rfc5424/parser.go:5209
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr209
		}
		goto st0
tr209:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st197
	st197:
		if p++; p == pe {
			goto _test_eof197
		}
	st_case_197:
//line rfc5424/parser.go:5228
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr210
		}
		goto st0
tr210:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st198
	st198:
		if p++; p == pe {
			goto _test_eof198
		}
	st_case_198:
//line rfc5424/parser.go:5247
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr211
		}
		goto st0
tr211:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st199
	st199:
		if p++; p == pe {
			goto _test_eof199
		}
	st_case_199:
//line rfc5424/parser.go:5266
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr212
		}
		goto st0
tr212:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st200
	st200:
		if p++; p == pe {
			goto _test_eof200
		}
	st_case_200:
//line rfc5424/parser.go:5285
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr213
		}
		goto st0
tr213:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st201
	st201:
		if p++; p == pe {
			goto _test_eof201
		}
	st_case_201:
//line rfc5424/parser.go:5304
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr214
		}
		goto st0
tr214:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st202
	st202:
		if p++; p == pe {
			goto _test_eof202
		}
	st_case_202:
//line rfc5424/parser.go:5323
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr215
		}
		goto st0
tr215:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st203
	st203:
		if p++; p == pe {
			goto _test_eof203
		}
	st_case_203:
//line rfc5424/parser.go:5342
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr216
		}
		goto st0
tr216:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st204
	st204:
		if p++; p == pe {
			goto _test_eof204
		}
	st_case_204:
//line rfc5424/parser.go:5361
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr217
		}
		goto st0
tr217:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st205
	st205:
		if p++; p == pe {
			goto _test_eof205
		}
	st_case_205:
//line rfc5424/parser.go:5380
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr218
		}
		goto st0
tr218:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st206
	st206:
		if p++; p == pe {
			goto _test_eof206
		}
	st_case_206:
//line rfc5424/parser.go:5399
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr219
		}
		goto st0
tr219:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st207
	st207:
		if p++; p == pe {
			goto _test_eof207
		}
	st_case_207:
//line rfc5424/parser.go:5418
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr220
		}
		goto st0
tr220:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st208
	st208:
		if p++; p == pe {
			goto _test_eof208
		}
	st_case_208:
//line rfc5424/parser.go:5437
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr221
		}
		goto st0
tr221:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st209
	st209:
		if p++; p == pe {
			goto _test_eof209
		}
	st_case_209:
//line rfc5424/parser.go:5456
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr222
		}
		goto st0
tr222:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st210
	st210:
		if p++; p == pe {
			goto _test_eof210
		}
	st_case_210:
//line rfc5424/parser.go:5475
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr223
		}
		goto st0
tr223:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st211
	st211:
		if p++; p == pe {
			goto _test_eof211
		}
	st_case_211:
//line rfc5424/parser.go:5494
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr224
		}
		goto st0
tr224:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st212
	st212:
		if p++; p == pe {
			goto _test_eof212
		}
	st_case_212:
//line rfc5424/parser.go:5513
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr225
		}
		goto st0
tr225:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st213
	st213:
		if p++; p == pe {
			goto _test_eof213
		}
	st_case_213:
//line rfc5424/parser.go:5532
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr226
		}
		goto st0
tr226:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st214
	st214:
		if p++; p == pe {
			goto _test_eof214
		}
	st_case_214:
//line rfc5424/parser.go:5551
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr227
		}
		goto st0
tr227:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st215
	st215:
		if p++; p == pe {
			goto _test_eof215
		}
	st_case_215:
//line rfc5424/parser.go:5570
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr228
		}
		goto st0
tr228:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st216
	st216:
		if p++; p == pe {
			goto _test_eof216
		}
	st_case_216:
//line rfc5424/parser.go:5589
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr229
		}
		goto st0
tr229:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st217
	st217:
		if p++; p == pe {
			goto _test_eof217
		}
	st_case_217:
//line rfc5424/parser.go:5608
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr230
		}
		goto st0
tr230:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st218
	st218:
		if p++; p == pe {
			goto _test_eof218
		}
	st_case_218:
//line rfc5424/parser.go:5627
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr231
		}
		goto st0
tr231:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st219
	st219:
		if p++; p == pe {
			goto _test_eof219
		}
	st_case_219:
//line rfc5424/parser.go:5646
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr232
		}
		goto st0
tr232:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st220
	st220:
		if p++; p == pe {
			goto _test_eof220
		}
	st_case_220:
//line rfc5424/parser.go:5665
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr233
		}
		goto st0
tr233:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st221
	st221:
		if p++; p == pe {
			goto _test_eof221
		}
	st_case_221:
//line rfc5424/parser.go:5684
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr234
		}
		goto st0
tr234:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st222
	st222:
		if p++; p == pe {
			goto _test_eof222
		}
	st_case_222:
//line rfc5424/parser.go:5703
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr235
		}
		goto st0
tr235:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st223
	st223:
		if p++; p == pe {
			goto _test_eof223
		}
	st_case_223:
//line rfc5424/parser.go:5722
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr236
		}
		goto st0
tr236:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st224
	st224:
		if p++; p == pe {
			goto _test_eof224
		}
	st_case_224:
//line rfc5424/parser.go:5741
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr237
		}
		goto st0
tr237:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st225
	st225:
		if p++; p == pe {
			goto _test_eof225
		}
	st_case_225:
//line rfc5424/parser.go:5760
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr238
		}
		goto st0
tr238:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st226
	st226:
		if p++; p == pe {
			goto _test_eof226
		}
	st_case_226:
//line rfc5424/parser.go:5779
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr239
		}
		goto st0
tr239:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st227
	st227:
		if p++; p == pe {
			goto _test_eof227
		}
	st_case_227:
//line rfc5424/parser.go:5798
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr240
		}
		goto st0
tr240:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st228
	st228:
		if p++; p == pe {
			goto _test_eof228
		}
	st_case_228:
//line rfc5424/parser.go:5817
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr241
		}
		goto st0
tr241:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st229
	st229:
		if p++; p == pe {
			goto _test_eof229
		}
	st_case_229:
//line rfc5424/parser.go:5836
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr242
		}
		goto st0
tr242:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st230
	st230:
		if p++; p == pe {
			goto _test_eof230
		}
	st_case_230:
//line rfc5424/parser.go:5855
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr243
		}
		goto st0
tr243:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st231
	st231:
		if p++; p == pe {
			goto _test_eof231
		}
	st_case_231:
//line rfc5424/parser.go:5874
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr244
		}
		goto st0
tr244:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st232
	st232:
		if p++; p == pe {
			goto _test_eof232
		}
	st_case_232:
//line rfc5424/parser.go:5893
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr245
		}
		goto st0
tr245:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st233
	st233:
		if p++; p == pe {
			goto _test_eof233
		}
	st_case_233:
//line rfc5424/parser.go:5912
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr246
		}
		goto st0
tr246:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st234
	st234:
		if p++; p == pe {
			goto _test_eof234
		}
	st_case_234:
//line rfc5424/parser.go:5931
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr247
		}
		goto st0
tr247:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st235
	st235:
		if p++; p == pe {
			goto _test_eof235
		}
	st_case_235:
//line rfc5424/parser.go:5950
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr248
		}
		goto st0
tr248:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st236
	st236:
		if p++; p == pe {
			goto _test_eof236
		}
	st_case_236:
//line rfc5424/parser.go:5969
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr249
		}
		goto st0
tr249:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st237
	st237:
		if p++; p == pe {
			goto _test_eof237
		}
	st_case_237:
//line rfc5424/parser.go:5988
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr250
		}
		goto st0
tr250:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st238
	st238:
		if p++; p == pe {
			goto _test_eof238
		}
	st_case_238:
//line rfc5424/parser.go:6007
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr251
		}
		goto st0
tr251:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st239
	st239:
		if p++; p == pe {
			goto _test_eof239
		}
	st_case_239:
//line rfc5424/parser.go:6026
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr252
		}
		goto st0
tr252:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st240
	st240:
		if p++; p == pe {
			goto _test_eof240
		}
	st_case_240:
//line rfc5424/parser.go:6045
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr253
		}
		goto st0
tr253:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st241
	st241:
		if p++; p == pe {
			goto _test_eof241
		}
	st_case_241:
//line rfc5424/parser.go:6064
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr254
		}
		goto st0
tr254:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st242
	st242:
		if p++; p == pe {
			goto _test_eof242
		}
	st_case_242:
//line rfc5424/parser.go:6083
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr255
		}
		goto st0
tr255:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st243
	st243:
		if p++; p == pe {
			goto _test_eof243
		}
	st_case_243:
//line rfc5424/parser.go:6102
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr256
		}
		goto st0
tr256:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st244
	st244:
		if p++; p == pe {
			goto _test_eof244
		}
	st_case_244:
//line rfc5424/parser.go:6121
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr257
		}
		goto st0
tr257:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st245
	st245:
		if p++; p == pe {
			goto _test_eof245
		}
	st_case_245:
//line rfc5424/parser.go:6140
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr258
		}
		goto st0
tr258:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st246
	st246:
		if p++; p == pe {
			goto _test_eof246
		}
	st_case_246:
//line rfc5424/parser.go:6159
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr259
		}
		goto st0
tr259:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st247
	st247:
		if p++; p == pe {
			goto _test_eof247
		}
	st_case_247:
//line rfc5424/parser.go:6178
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr260
		}
		goto st0
tr260:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st248
	st248:
		if p++; p == pe {
			goto _test_eof248
		}
	st_case_248:
//line rfc5424/parser.go:6197
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr261
		}
		goto st0
tr261:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st249
	st249:
		if p++; p == pe {
			goto _test_eof249
		}
	st_case_249:
//line rfc5424/parser.go:6216
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr262
		}
		goto st0
tr262:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st250
	st250:
		if p++; p == pe {
			goto _test_eof250
		}
	st_case_250:
//line rfc5424/parser.go:6235
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr263
		}
		goto st0
tr263:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st251
	st251:
		if p++; p == pe {
			goto _test_eof251
		}
	st_case_251:
//line rfc5424/parser.go:6254
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr264
		}
		goto st0
tr264:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st252
	st252:
		if p++; p == pe {
			goto _test_eof252
		}
	st_case_252:
//line rfc5424/parser.go:6273
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr265
		}
		goto st0
tr265:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st253
	st253:
		if p++; p == pe {
			goto _test_eof253
		}
	st_case_253:
//line rfc5424/parser.go:6292
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr266
		}
		goto st0
tr266:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st254
	st254:
		if p++; p == pe {
			goto _test_eof254
		}
	st_case_254:
//line rfc5424/parser.go:6311
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr267
		}
		goto st0
tr267:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st255
	st255:
		if p++; p == pe {
			goto _test_eof255
		}
	st_case_255:
//line rfc5424/parser.go:6330
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr268
		}
		goto st0
tr268:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st256
	st256:
		if p++; p == pe {
			goto _test_eof256
		}
	st_case_256:
//line rfc5424/parser.go:6349
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr269
		}
		goto st0
tr269:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st257
	st257:
		if p++; p == pe {
			goto _test_eof257
		}
	st_case_257:
//line rfc5424/parser.go:6368
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr270
		}
		goto st0
tr270:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st258
	st258:
		if p++; p == pe {
			goto _test_eof258
		}
	st_case_258:
//line rfc5424/parser.go:6387
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr271
		}
		goto st0
tr271:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st259
	st259:
		if p++; p == pe {
			goto _test_eof259
		}
	st_case_259:
//line rfc5424/parser.go:6406
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr272
		}
		goto st0
tr272:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st260
	st260:
		if p++; p == pe {
			goto _test_eof260
		}
	st_case_260:
//line rfc5424/parser.go:6425
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr273
		}
		goto st0
tr273:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st261
	st261:
		if p++; p == pe {
			goto _test_eof261
		}
	st_case_261:
//line rfc5424/parser.go:6444
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr274
		}
		goto st0
tr274:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st262
	st262:
		if p++; p == pe {
			goto _test_eof262
		}
	st_case_262:
//line rfc5424/parser.go:6463
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr275
		}
		goto st0
tr275:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st263
	st263:
		if p++; p == pe {
			goto _test_eof263
		}
	st_case_263:
//line rfc5424/parser.go:6482
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr276
		}
		goto st0
tr276:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st264
	st264:
		if p++; p == pe {
			goto _test_eof264
		}
	st_case_264:
//line rfc5424/parser.go:6501
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr277
		}
		goto st0
tr277:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st265
	st265:
		if p++; p == pe {
			goto _test_eof265
		}
	st_case_265:
//line rfc5424/parser.go:6520
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr278
		}
		goto st0
tr278:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st266
	st266:
		if p++; p == pe {
			goto _test_eof266
		}
	st_case_266:
//line rfc5424/parser.go:6539
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr279
		}
		goto st0
tr279:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st267
	st267:
		if p++; p == pe {
			goto _test_eof267
		}
	st_case_267:
//line rfc5424/parser.go:6558
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr280
		}
		goto st0
tr280:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st268
	st268:
		if p++; p == pe {
			goto _test_eof268
		}
	st_case_268:
//line rfc5424/parser.go:6577
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr281
		}
		goto st0
tr281:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st269
	st269:
		if p++; p == pe {
			goto _test_eof269
		}
	st_case_269:
//line rfc5424/parser.go:6596
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr282
		}
		goto st0
tr282:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st270
	st270:
		if p++; p == pe {
			goto _test_eof270
		}
	st_case_270:
//line rfc5424/parser.go:6615
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr283
		}
		goto st0
tr283:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st271
	st271:
		if p++; p == pe {
			goto _test_eof271
		}
	st_case_271:
//line rfc5424/parser.go:6634
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr284
		}
		goto st0
tr284:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st272
	st272:
		if p++; p == pe {
			goto _test_eof272
		}
	st_case_272:
//line rfc5424/parser.go:6653
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr285
		}
		goto st0
tr285:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st273
	st273:
		if p++; p == pe {
			goto _test_eof273
		}
	st_case_273:
//line rfc5424/parser.go:6672
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr286
		}
		goto st0
tr286:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st274
	st274:
		if p++; p == pe {
			goto _test_eof274
		}
	st_case_274:
//line rfc5424/parser.go:6691
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr287
		}
		goto st0
tr287:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st275
	st275:
		if p++; p == pe {
			goto _test_eof275
		}
	st_case_275:
//line rfc5424/parser.go:6710
		if data[p] == 32 {
			goto tr23
		}
		goto st0
tr20:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st276
	st276:
		if p++; p == pe {
			goto _test_eof276
		}
	st_case_276:
//line rfc5424/parser.go:6726
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr288
		}
		goto st0
tr288:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st277
	st277:
		if p++; p == pe {
			goto _test_eof277
		}
	st_case_277:
//line rfc5424/parser.go:6745
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr289
		}
		goto st0
tr289:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st278
	st278:
		if p++; p == pe {
			goto _test_eof278
		}
	st_case_278:
//line rfc5424/parser.go:6764
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr290
		}
		goto st0
tr290:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st279
	st279:
		if p++; p == pe {
			goto _test_eof279
		}
	st_case_279:
//line rfc5424/parser.go:6783
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr291
		}
		goto st0
tr291:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st280
	st280:
		if p++; p == pe {
			goto _test_eof280
		}
	st_case_280:
//line rfc5424/parser.go:6802
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr292
		}
		goto st0
tr292:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st281
	st281:
		if p++; p == pe {
			goto _test_eof281
		}
	st_case_281:
//line rfc5424/parser.go:6821
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr293
		}
		goto st0
tr293:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st282
	st282:
		if p++; p == pe {
			goto _test_eof282
		}
	st_case_282:
//line rfc5424/parser.go:6840
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr294
		}
		goto st0
tr294:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st283
	st283:
		if p++; p == pe {
			goto _test_eof283
		}
	st_case_283:
//line rfc5424/parser.go:6859
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr295
		}
		goto st0
tr295:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st284
	st284:
		if p++; p == pe {
			goto _test_eof284
		}
	st_case_284:
//line rfc5424/parser.go:6878
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr296
		}
		goto st0
tr296:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st285
	st285:
		if p++; p == pe {
			goto _test_eof285
		}
	st_case_285:
//line rfc5424/parser.go:6897
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr297
		}
		goto st0
tr297:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st286
	st286:
		if p++; p == pe {
			goto _test_eof286
		}
	st_case_286:
//line rfc5424/parser.go:6916
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr298
		}
		goto st0
tr298:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st287
	st287:
		if p++; p == pe {
			goto _test_eof287
		}
	st_case_287:
//line rfc5424/parser.go:6935
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr299
		}
		goto st0
tr299:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st288
	st288:
		if p++; p == pe {
			goto _test_eof288
		}
	st_case_288:
//line rfc5424/parser.go:6954
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr300
		}
		goto st0
tr300:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st289
	st289:
		if p++; p == pe {
			goto _test_eof289
		}
	st_case_289:
//line rfc5424/parser.go:6973
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr301
		}
		goto st0
tr301:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st290
	st290:
		if p++; p == pe {
			goto _test_eof290
		}
	st_case_290:
//line rfc5424/parser.go:6992
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr302
		}
		goto st0
tr302:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st291
	st291:
		if p++; p == pe {
			goto _test_eof291
		}
	st_case_291:
//line rfc5424/parser.go:7011
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr303
		}
		goto st0
tr303:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st292
	st292:
		if p++; p == pe {
			goto _test_eof292
		}
	st_case_292:
//line rfc5424/parser.go:7030
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr304
		}
		goto st0
tr304:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st293
	st293:
		if p++; p == pe {
			goto _test_eof293
		}
	st_case_293:
//line rfc5424/parser.go:7049
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr305
		}
		goto st0
tr305:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st294
	st294:
		if p++; p == pe {
			goto _test_eof294
		}
	st_case_294:
//line rfc5424/parser.go:7068
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr306
		}
		goto st0
tr306:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st295
	st295:
		if p++; p == pe {
			goto _test_eof295
		}
	st_case_295:
//line rfc5424/parser.go:7087
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr307
		}
		goto st0
tr307:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st296
	st296:
		if p++; p == pe {
			goto _test_eof296
		}
	st_case_296:
//line rfc5424/parser.go:7106
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr308
		}
		goto st0
tr308:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st297
	st297:
		if p++; p == pe {
			goto _test_eof297
		}
	st_case_297:
//line rfc5424/parser.go:7125
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr309
		}
		goto st0
tr309:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st298
	st298:
		if p++; p == pe {
			goto _test_eof298
		}
	st_case_298:
//line rfc5424/parser.go:7144
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr310
		}
		goto st0
tr310:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st299
	st299:
		if p++; p == pe {
			goto _test_eof299
		}
	st_case_299:
//line rfc5424/parser.go:7163
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr311
		}
		goto st0
tr311:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st300
	st300:
		if p++; p == pe {
			goto _test_eof300
		}
	st_case_300:
//line rfc5424/parser.go:7182
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr312
		}
		goto st0
tr312:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st301
	st301:
		if p++; p == pe {
			goto _test_eof301
		}
	st_case_301:
//line rfc5424/parser.go:7201
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr313
		}
		goto st0
tr313:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st302
	st302:
		if p++; p == pe {
			goto _test_eof302
		}
	st_case_302:
//line rfc5424/parser.go:7220
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr314
		}
		goto st0
tr314:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st303
	st303:
		if p++; p == pe {
			goto _test_eof303
		}
	st_case_303:
//line rfc5424/parser.go:7239
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr315
		}
		goto st0
tr315:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st304
	st304:
		if p++; p == pe {
			goto _test_eof304
		}
	st_case_304:
//line rfc5424/parser.go:7258
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr316
		}
		goto st0
tr316:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st305
	st305:
		if p++; p == pe {
			goto _test_eof305
		}
	st_case_305:
//line rfc5424/parser.go:7277
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr317
		}
		goto st0
tr317:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st306
	st306:
		if p++; p == pe {
			goto _test_eof306
		}
	st_case_306:
//line rfc5424/parser.go:7296
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr318
		}
		goto st0
tr318:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st307
	st307:
		if p++; p == pe {
			goto _test_eof307
		}
	st_case_307:
//line rfc5424/parser.go:7315
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr319
		}
		goto st0
tr319:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st308
	st308:
		if p++; p == pe {
			goto _test_eof308
		}
	st_case_308:
//line rfc5424/parser.go:7334
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr320
		}
		goto st0
tr320:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st309
	st309:
		if p++; p == pe {
			goto _test_eof309
		}
	st_case_309:
//line rfc5424/parser.go:7353
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr321
		}
		goto st0
tr321:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st310
	st310:
		if p++; p == pe {
			goto _test_eof310
		}
	st_case_310:
//line rfc5424/parser.go:7372
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr322
		}
		goto st0
tr322:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st311
	st311:
		if p++; p == pe {
			goto _test_eof311
		}
	st_case_311:
//line rfc5424/parser.go:7391
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr323
		}
		goto st0
tr323:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st312
	st312:
		if p++; p == pe {
			goto _test_eof312
		}
	st_case_312:
//line rfc5424/parser.go:7410
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr324
		}
		goto st0
tr324:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st313
	st313:
		if p++; p == pe {
			goto _test_eof313
		}
	st_case_313:
//line rfc5424/parser.go:7429
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr325
		}
		goto st0
tr325:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st314
	st314:
		if p++; p == pe {
			goto _test_eof314
		}
	st_case_314:
//line rfc5424/parser.go:7448
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr326
		}
		goto st0
tr326:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st315
	st315:
		if p++; p == pe {
			goto _test_eof315
		}
	st_case_315:
//line rfc5424/parser.go:7467
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr327
		}
		goto st0
tr327:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st316
	st316:
		if p++; p == pe {
			goto _test_eof316
		}
	st_case_316:
//line rfc5424/parser.go:7486
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr328
		}
		goto st0
tr328:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st317
	st317:
		if p++; p == pe {
			goto _test_eof317
		}
	st_case_317:
//line rfc5424/parser.go:7505
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr329
		}
		goto st0
tr329:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st318
	st318:
		if p++; p == pe {
			goto _test_eof318
		}
	st_case_318:
//line rfc5424/parser.go:7524
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr330
		}
		goto st0
tr330:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st319
	st319:
		if p++; p == pe {
			goto _test_eof319
		}
	st_case_319:
//line rfc5424/parser.go:7543
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr331
		}
		goto st0
tr331:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st320
	st320:
		if p++; p == pe {
			goto _test_eof320
		}
	st_case_320:
//line rfc5424/parser.go:7562
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr332
		}
		goto st0
tr332:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st321
	st321:
		if p++; p == pe {
			goto _test_eof321
		}
	st_case_321:
//line rfc5424/parser.go:7581
		if data[p] == 32 {
			goto tr19
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr333
		}
		goto st0
tr333:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st322
	st322:
		if p++; p == pe {
			goto _test_eof322
		}
	st_case_322:
//line rfc5424/parser.go:7600
		if data[p] == 32 {
			goto tr19
		}
		goto st0
tr16:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st323
	st323:
		if p++; p == pe {
			goto _test_eof323
		}
	st_case_323:
//line rfc5424/parser.go:7616
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr334
		}
		goto st0
tr334:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st324
	st324:
		if p++; p == pe {
			goto _test_eof324
		}
	st_case_324:
//line rfc5424/parser.go:7635
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr335
		}
		goto st0
tr335:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st325
	st325:
		if p++; p == pe {
			goto _test_eof325
		}
	st_case_325:
//line rfc5424/parser.go:7654
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr336
		}
		goto st0
tr336:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st326
	st326:
		if p++; p == pe {
			goto _test_eof326
		}
	st_case_326:
//line rfc5424/parser.go:7673
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr337
		}
		goto st0
tr337:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st327
	st327:
		if p++; p == pe {
			goto _test_eof327
		}
	st_case_327:
//line rfc5424/parser.go:7692
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr338
		}
		goto st0
tr338:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st328
	st328:
		if p++; p == pe {
			goto _test_eof328
		}
	st_case_328:
//line rfc5424/parser.go:7711
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr339
		}
		goto st0
tr339:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st329
	st329:
		if p++; p == pe {
			goto _test_eof329
		}
	st_case_329:
//line rfc5424/parser.go:7730
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr340
		}
		goto st0
tr340:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st330
	st330:
		if p++; p == pe {
			goto _test_eof330
		}
	st_case_330:
//line rfc5424/parser.go:7749
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr341
		}
		goto st0
tr341:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st331
	st331:
		if p++; p == pe {
			goto _test_eof331
		}
	st_case_331:
//line rfc5424/parser.go:7768
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr342
		}
		goto st0
tr342:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st332
	st332:
		if p++; p == pe {
			goto _test_eof332
		}
	st_case_332:
//line rfc5424/parser.go:7787
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr343
		}
		goto st0
tr343:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st333
	st333:
		if p++; p == pe {
			goto _test_eof333
		}
	st_case_333:
//line rfc5424/parser.go:7806
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr344
		}
		goto st0
tr344:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st334
	st334:
		if p++; p == pe {
			goto _test_eof334
		}
	st_case_334:
//line rfc5424/parser.go:7825
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr345
		}
		goto st0
tr345:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st335
	st335:
		if p++; p == pe {
			goto _test_eof335
		}
	st_case_335:
//line rfc5424/parser.go:7844
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr346
		}
		goto st0
tr346:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st336
	st336:
		if p++; p == pe {
			goto _test_eof336
		}
	st_case_336:
//line rfc5424/parser.go:7863
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr347
		}
		goto st0
tr347:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st337
	st337:
		if p++; p == pe {
			goto _test_eof337
		}
	st_case_337:
//line rfc5424/parser.go:7882
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr348
		}
		goto st0
tr348:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st338
	st338:
		if p++; p == pe {
			goto _test_eof338
		}
	st_case_338:
//line rfc5424/parser.go:7901
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr349
		}
		goto st0
tr349:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st339
	st339:
		if p++; p == pe {
			goto _test_eof339
		}
	st_case_339:
//line rfc5424/parser.go:7920
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr350
		}
		goto st0
tr350:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st340
	st340:
		if p++; p == pe {
			goto _test_eof340
		}
	st_case_340:
//line rfc5424/parser.go:7939
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr351
		}
		goto st0
tr351:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st341
	st341:
		if p++; p == pe {
			goto _test_eof341
		}
	st_case_341:
//line rfc5424/parser.go:7958
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr352
		}
		goto st0
tr352:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st342
	st342:
		if p++; p == pe {
			goto _test_eof342
		}
	st_case_342:
//line rfc5424/parser.go:7977
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr353
		}
		goto st0
tr353:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st343
	st343:
		if p++; p == pe {
			goto _test_eof343
		}
	st_case_343:
//line rfc5424/parser.go:7996
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr354
		}
		goto st0
tr354:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st344
	st344:
		if p++; p == pe {
			goto _test_eof344
		}
	st_case_344:
//line rfc5424/parser.go:8015
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr355
		}
		goto st0
tr355:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st345
	st345:
		if p++; p == pe {
			goto _test_eof345
		}
	st_case_345:
//line rfc5424/parser.go:8034
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr356
		}
		goto st0
tr356:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st346
	st346:
		if p++; p == pe {
			goto _test_eof346
		}
	st_case_346:
//line rfc5424/parser.go:8053
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr357
		}
		goto st0
tr357:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st347
	st347:
		if p++; p == pe {
			goto _test_eof347
		}
	st_case_347:
//line rfc5424/parser.go:8072
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr358
		}
		goto st0
tr358:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st348
	st348:
		if p++; p == pe {
			goto _test_eof348
		}
	st_case_348:
//line rfc5424/parser.go:8091
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr359
		}
		goto st0
tr359:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st349
	st349:
		if p++; p == pe {
			goto _test_eof349
		}
	st_case_349:
//line rfc5424/parser.go:8110
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr360
		}
		goto st0
tr360:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st350
	st350:
		if p++; p == pe {
			goto _test_eof350
		}
	st_case_350:
//line rfc5424/parser.go:8129
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr361
		}
		goto st0
tr361:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st351
	st351:
		if p++; p == pe {
			goto _test_eof351
		}
	st_case_351:
//line rfc5424/parser.go:8148
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr362
		}
		goto st0
tr362:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st352
	st352:
		if p++; p == pe {
			goto _test_eof352
		}
	st_case_352:
//line rfc5424/parser.go:8167
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr363
		}
		goto st0
tr363:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st353
	st353:
		if p++; p == pe {
			goto _test_eof353
		}
	st_case_353:
//line rfc5424/parser.go:8186
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr364
		}
		goto st0
tr364:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st354
	st354:
		if p++; p == pe {
			goto _test_eof354
		}
	st_case_354:
//line rfc5424/parser.go:8205
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr365
		}
		goto st0
tr365:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st355
	st355:
		if p++; p == pe {
			goto _test_eof355
		}
	st_case_355:
//line rfc5424/parser.go:8224
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr366
		}
		goto st0
tr366:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st356
	st356:
		if p++; p == pe {
			goto _test_eof356
		}
	st_case_356:
//line rfc5424/parser.go:8243
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr367
		}
		goto st0
tr367:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st357
	st357:
		if p++; p == pe {
			goto _test_eof357
		}
	st_case_357:
//line rfc5424/parser.go:8262
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr368
		}
		goto st0
tr368:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st358
	st358:
		if p++; p == pe {
			goto _test_eof358
		}
	st_case_358:
//line rfc5424/parser.go:8281
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr369
		}
		goto st0
tr369:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st359
	st359:
		if p++; p == pe {
			goto _test_eof359
		}
	st_case_359:
//line rfc5424/parser.go:8300
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr370
		}
		goto st0
tr370:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st360
	st360:
		if p++; p == pe {
			goto _test_eof360
		}
	st_case_360:
//line rfc5424/parser.go:8319
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr371
		}
		goto st0
tr371:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st361
	st361:
		if p++; p == pe {
			goto _test_eof361
		}
	st_case_361:
//line rfc5424/parser.go:8338
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr372
		}
		goto st0
tr372:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st362
	st362:
		if p++; p == pe {
			goto _test_eof362
		}
	st_case_362:
//line rfc5424/parser.go:8357
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr373
		}
		goto st0
tr373:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st363
	st363:
		if p++; p == pe {
			goto _test_eof363
		}
	st_case_363:
//line rfc5424/parser.go:8376
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr374
		}
		goto st0
tr374:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st364
	st364:
		if p++; p == pe {
			goto _test_eof364
		}
	st_case_364:
//line rfc5424/parser.go:8395
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr375
		}
		goto st0
tr375:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st365
	st365:
		if p++; p == pe {
			goto _test_eof365
		}
	st_case_365:
//line rfc5424/parser.go:8414
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr376
		}
		goto st0
tr376:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st366
	st366:
		if p++; p == pe {
			goto _test_eof366
		}
	st_case_366:
//line rfc5424/parser.go:8433
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr377
		}
		goto st0
tr377:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st367
	st367:
		if p++; p == pe {
			goto _test_eof367
		}
	st_case_367:
//line rfc5424/parser.go:8452
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr378
		}
		goto st0
tr378:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st368
	st368:
		if p++; p == pe {
			goto _test_eof368
		}
	st_case_368:
//line rfc5424/parser.go:8471
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr379
		}
		goto st0
tr379:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st369
	st369:
		if p++; p == pe {
			goto _test_eof369
		}
	st_case_369:
//line rfc5424/parser.go:8490
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr380
		}
		goto st0
tr380:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st370
	st370:
		if p++; p == pe {
			goto _test_eof370
		}
	st_case_370:
//line rfc5424/parser.go:8509
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr381
		}
		goto st0
tr381:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st371
	st371:
		if p++; p == pe {
			goto _test_eof371
		}
	st_case_371:
//line rfc5424/parser.go:8528
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr382
		}
		goto st0
tr382:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st372
	st372:
		if p++; p == pe {
			goto _test_eof372
		}
	st_case_372:
//line rfc5424/parser.go:8547
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr383
		}
		goto st0
tr383:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st373
	st373:
		if p++; p == pe {
			goto _test_eof373
		}
	st_case_373:
//line rfc5424/parser.go:8566
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr384
		}
		goto st0
tr384:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st374
	st374:
		if p++; p == pe {
			goto _test_eof374
		}
	st_case_374:
//line rfc5424/parser.go:8585
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr385
		}
		goto st0
tr385:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st375
	st375:
		if p++; p == pe {
			goto _test_eof375
		}
	st_case_375:
//line rfc5424/parser.go:8604
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr386
		}
		goto st0
tr386:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st376
	st376:
		if p++; p == pe {
			goto _test_eof376
		}
	st_case_376:
//line rfc5424/parser.go:8623
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr387
		}
		goto st0
tr387:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st377
	st377:
		if p++; p == pe {
			goto _test_eof377
		}
	st_case_377:
//line rfc5424/parser.go:8642
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr388
		}
		goto st0
tr388:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st378
	st378:
		if p++; p == pe {
			goto _test_eof378
		}
	st_case_378:
//line rfc5424/parser.go:8661
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr389
		}
		goto st0
tr389:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st379
	st379:
		if p++; p == pe {
			goto _test_eof379
		}
	st_case_379:
//line rfc5424/parser.go:8680
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr390
		}
		goto st0
tr390:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st380
	st380:
		if p++; p == pe {
			goto _test_eof380
		}
	st_case_380:
//line rfc5424/parser.go:8699
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr391
		}
		goto st0
tr391:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st381
	st381:
		if p++; p == pe {
			goto _test_eof381
		}
	st_case_381:
//line rfc5424/parser.go:8718
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr392
		}
		goto st0
tr392:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st382
	st382:
		if p++; p == pe {
			goto _test_eof382
		}
	st_case_382:
//line rfc5424/parser.go:8737
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr393
		}
		goto st0
tr393:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st383
	st383:
		if p++; p == pe {
			goto _test_eof383
		}
	st_case_383:
//line rfc5424/parser.go:8756
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr394
		}
		goto st0
tr394:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st384
	st384:
		if p++; p == pe {
			goto _test_eof384
		}
	st_case_384:
//line rfc5424/parser.go:8775
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr395
		}
		goto st0
tr395:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st385
	st385:
		if p++; p == pe {
			goto _test_eof385
		}
	st_case_385:
//line rfc5424/parser.go:8794
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr396
		}
		goto st0
tr396:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st386
	st386:
		if p++; p == pe {
			goto _test_eof386
		}
	st_case_386:
//line rfc5424/parser.go:8813
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr397
		}
		goto st0
tr397:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st387
	st387:
		if p++; p == pe {
			goto _test_eof387
		}
	st_case_387:
//line rfc5424/parser.go:8832
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr398
		}
		goto st0
tr398:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st388
	st388:
		if p++; p == pe {
			goto _test_eof388
		}
	st_case_388:
//line rfc5424/parser.go:8851
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr399
		}
		goto st0
tr399:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st389
	st389:
		if p++; p == pe {
			goto _test_eof389
		}
	st_case_389:
//line rfc5424/parser.go:8870
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr400
		}
		goto st0
tr400:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st390
	st390:
		if p++; p == pe {
			goto _test_eof390
		}
	st_case_390:
//line rfc5424/parser.go:8889
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr401
		}
		goto st0
tr401:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st391
	st391:
		if p++; p == pe {
			goto _test_eof391
		}
	st_case_391:
//line rfc5424/parser.go:8908
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr402
		}
		goto st0
tr402:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st392
	st392:
		if p++; p == pe {
			goto _test_eof392
		}
	st_case_392:
//line rfc5424/parser.go:8927
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr403
		}
		goto st0
tr403:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st393
	st393:
		if p++; p == pe {
			goto _test_eof393
		}
	st_case_393:
//line rfc5424/parser.go:8946
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr404
		}
		goto st0
tr404:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st394
	st394:
		if p++; p == pe {
			goto _test_eof394
		}
	st_case_394:
//line rfc5424/parser.go:8965
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr405
		}
		goto st0
tr405:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st395
	st395:
		if p++; p == pe {
			goto _test_eof395
		}
	st_case_395:
//line rfc5424/parser.go:8984
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr406
		}
		goto st0
tr406:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st396
	st396:
		if p++; p == pe {
			goto _test_eof396
		}
	st_case_396:
//line rfc5424/parser.go:9003
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr407
		}
		goto st0
tr407:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st397
	st397:
		if p++; p == pe {
			goto _test_eof397
		}
	st_case_397:
//line rfc5424/parser.go:9022
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr408
		}
		goto st0
tr408:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st398
	st398:
		if p++; p == pe {
			goto _test_eof398
		}
	st_case_398:
//line rfc5424/parser.go:9041
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr409
		}
		goto st0
tr409:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st399
	st399:
		if p++; p == pe {
			goto _test_eof399
		}
	st_case_399:
//line rfc5424/parser.go:9060
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr410
		}
		goto st0
tr410:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st400
	st400:
		if p++; p == pe {
			goto _test_eof400
		}
	st_case_400:
//line rfc5424/parser.go:9079
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr411
		}
		goto st0
tr411:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st401
	st401:
		if p++; p == pe {
			goto _test_eof401
		}
	st_case_401:
//line rfc5424/parser.go:9098
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr412
		}
		goto st0
tr412:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st402
	st402:
		if p++; p == pe {
			goto _test_eof402
		}
	st_case_402:
//line rfc5424/parser.go:9117
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr413
		}
		goto st0
tr413:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st403
	st403:
		if p++; p == pe {
			goto _test_eof403
		}
	st_case_403:
//line rfc5424/parser.go:9136
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr414
		}
		goto st0
tr414:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st404
	st404:
		if p++; p == pe {
			goto _test_eof404
		}
	st_case_404:
//line rfc5424/parser.go:9155
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr415
		}
		goto st0
tr415:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st405
	st405:
		if p++; p == pe {
			goto _test_eof405
		}
	st_case_405:
//line rfc5424/parser.go:9174
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr416
		}
		goto st0
tr416:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st406
	st406:
		if p++; p == pe {
			goto _test_eof406
		}
	st_case_406:
//line rfc5424/parser.go:9193
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr417
		}
		goto st0
tr417:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st407
	st407:
		if p++; p == pe {
			goto _test_eof407
		}
	st_case_407:
//line rfc5424/parser.go:9212
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr418
		}
		goto st0
tr418:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st408
	st408:
		if p++; p == pe {
			goto _test_eof408
		}
	st_case_408:
//line rfc5424/parser.go:9231
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr419
		}
		goto st0
tr419:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st409
	st409:
		if p++; p == pe {
			goto _test_eof409
		}
	st_case_409:
//line rfc5424/parser.go:9250
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr420
		}
		goto st0
tr420:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st410
	st410:
		if p++; p == pe {
			goto _test_eof410
		}
	st_case_410:
//line rfc5424/parser.go:9269
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr421
		}
		goto st0
tr421:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st411
	st411:
		if p++; p == pe {
			goto _test_eof411
		}
	st_case_411:
//line rfc5424/parser.go:9288
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr422
		}
		goto st0
tr422:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st412
	st412:
		if p++; p == pe {
			goto _test_eof412
		}
	st_case_412:
//line rfc5424/parser.go:9307
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr423
		}
		goto st0
tr423:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st413
	st413:
		if p++; p == pe {
			goto _test_eof413
		}
	st_case_413:
//line rfc5424/parser.go:9326
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr424
		}
		goto st0
tr424:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st414
	st414:
		if p++; p == pe {
			goto _test_eof414
		}
	st_case_414:
//line rfc5424/parser.go:9345
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr425
		}
		goto st0
tr425:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st415
	st415:
		if p++; p == pe {
			goto _test_eof415
		}
	st_case_415:
//line rfc5424/parser.go:9364
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr426
		}
		goto st0
tr426:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st416
	st416:
		if p++; p == pe {
			goto _test_eof416
		}
	st_case_416:
//line rfc5424/parser.go:9383
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr427
		}
		goto st0
tr427:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st417
	st417:
		if p++; p == pe {
			goto _test_eof417
		}
	st_case_417:
//line rfc5424/parser.go:9402
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr428
		}
		goto st0
tr428:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st418
	st418:
		if p++; p == pe {
			goto _test_eof418
		}
	st_case_418:
//line rfc5424/parser.go:9421
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr429
		}
		goto st0
tr429:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st419
	st419:
		if p++; p == pe {
			goto _test_eof419
		}
	st_case_419:
//line rfc5424/parser.go:9440
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr430
		}
		goto st0
tr430:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st420
	st420:
		if p++; p == pe {
			goto _test_eof420
		}
	st_case_420:
//line rfc5424/parser.go:9459
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr431
		}
		goto st0
tr431:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st421
	st421:
		if p++; p == pe {
			goto _test_eof421
		}
	st_case_421:
//line rfc5424/parser.go:9478
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr432
		}
		goto st0
tr432:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st422
	st422:
		if p++; p == pe {
			goto _test_eof422
		}
	st_case_422:
//line rfc5424/parser.go:9497
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr433
		}
		goto st0
tr433:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st423
	st423:
		if p++; p == pe {
			goto _test_eof423
		}
	st_case_423:
//line rfc5424/parser.go:9516
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr434
		}
		goto st0
tr434:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st424
	st424:
		if p++; p == pe {
			goto _test_eof424
		}
	st_case_424:
//line rfc5424/parser.go:9535
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr435
		}
		goto st0
tr435:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st425
	st425:
		if p++; p == pe {
			goto _test_eof425
		}
	st_case_425:
//line rfc5424/parser.go:9554
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr436
		}
		goto st0
tr436:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st426
	st426:
		if p++; p == pe {
			goto _test_eof426
		}
	st_case_426:
//line rfc5424/parser.go:9573
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr437
		}
		goto st0
tr437:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st427
	st427:
		if p++; p == pe {
			goto _test_eof427
		}
	st_case_427:
//line rfc5424/parser.go:9592
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr438
		}
		goto st0
tr438:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st428
	st428:
		if p++; p == pe {
			goto _test_eof428
		}
	st_case_428:
//line rfc5424/parser.go:9611
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr439
		}
		goto st0
tr439:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st429
	st429:
		if p++; p == pe {
			goto _test_eof429
		}
	st_case_429:
//line rfc5424/parser.go:9630
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr440
		}
		goto st0
tr440:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st430
	st430:
		if p++; p == pe {
			goto _test_eof430
		}
	st_case_430:
//line rfc5424/parser.go:9649
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr441
		}
		goto st0
tr441:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st431
	st431:
		if p++; p == pe {
			goto _test_eof431
		}
	st_case_431:
//line rfc5424/parser.go:9668
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr442
		}
		goto st0
tr442:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st432
	st432:
		if p++; p == pe {
			goto _test_eof432
		}
	st_case_432:
//line rfc5424/parser.go:9687
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr443
		}
		goto st0
tr443:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st433
	st433:
		if p++; p == pe {
			goto _test_eof433
		}
	st_case_433:
//line rfc5424/parser.go:9706
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr444
		}
		goto st0
tr444:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st434
	st434:
		if p++; p == pe {
			goto _test_eof434
		}
	st_case_434:
//line rfc5424/parser.go:9725
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr445
		}
		goto st0
tr445:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st435
	st435:
		if p++; p == pe {
			goto _test_eof435
		}
	st_case_435:
//line rfc5424/parser.go:9744
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr446
		}
		goto st0
tr446:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st436
	st436:
		if p++; p == pe {
			goto _test_eof436
		}
	st_case_436:
//line rfc5424/parser.go:9763
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr447
		}
		goto st0
tr447:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st437
	st437:
		if p++; p == pe {
			goto _test_eof437
		}
	st_case_437:
//line rfc5424/parser.go:9782
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr448
		}
		goto st0
tr448:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st438
	st438:
		if p++; p == pe {
			goto _test_eof438
		}
	st_case_438:
//line rfc5424/parser.go:9801
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr449
		}
		goto st0
tr449:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st439
	st439:
		if p++; p == pe {
			goto _test_eof439
		}
	st_case_439:
//line rfc5424/parser.go:9820
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr450
		}
		goto st0
tr450:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st440
	st440:
		if p++; p == pe {
			goto _test_eof440
		}
	st_case_440:
//line rfc5424/parser.go:9839
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr451
		}
		goto st0
tr451:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st441
	st441:
		if p++; p == pe {
			goto _test_eof441
		}
	st_case_441:
//line rfc5424/parser.go:9858
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr452
		}
		goto st0
tr452:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st442
	st442:
		if p++; p == pe {
			goto _test_eof442
		}
	st_case_442:
//line rfc5424/parser.go:9877
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr453
		}
		goto st0
tr453:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st443
	st443:
		if p++; p == pe {
			goto _test_eof443
		}
	st_case_443:
//line rfc5424/parser.go:9896
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr454
		}
		goto st0
tr454:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st444
	st444:
		if p++; p == pe {
			goto _test_eof444
		}
	st_case_444:
//line rfc5424/parser.go:9915
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr455
		}
		goto st0
tr455:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st445
	st445:
		if p++; p == pe {
			goto _test_eof445
		}
	st_case_445:
//line rfc5424/parser.go:9934
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr456
		}
		goto st0
tr456:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st446
	st446:
		if p++; p == pe {
			goto _test_eof446
		}
	st_case_446:
//line rfc5424/parser.go:9953
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr457
		}
		goto st0
tr457:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st447
	st447:
		if p++; p == pe {
			goto _test_eof447
		}
	st_case_447:
//line rfc5424/parser.go:9972
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr458
		}
		goto st0
tr458:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st448
	st448:
		if p++; p == pe {
			goto _test_eof448
		}
	st_case_448:
//line rfc5424/parser.go:9991
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr459
		}
		goto st0
tr459:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st449
	st449:
		if p++; p == pe {
			goto _test_eof449
		}
	st_case_449:
//line rfc5424/parser.go:10010
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr460
		}
		goto st0
tr460:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st450
	st450:
		if p++; p == pe {
			goto _test_eof450
		}
	st_case_450:
//line rfc5424/parser.go:10029
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr461
		}
		goto st0
tr461:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st451
	st451:
		if p++; p == pe {
			goto _test_eof451
		}
	st_case_451:
//line rfc5424/parser.go:10048
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr462
		}
		goto st0
tr462:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st452
	st452:
		if p++; p == pe {
			goto _test_eof452
		}
	st_case_452:
//line rfc5424/parser.go:10067
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr463
		}
		goto st0
tr463:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st453
	st453:
		if p++; p == pe {
			goto _test_eof453
		}
	st_case_453:
//line rfc5424/parser.go:10086
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr464
		}
		goto st0
tr464:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st454
	st454:
		if p++; p == pe {
			goto _test_eof454
		}
	st_case_454:
//line rfc5424/parser.go:10105
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr465
		}
		goto st0
tr465:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st455
	st455:
		if p++; p == pe {
			goto _test_eof455
		}
	st_case_455:
//line rfc5424/parser.go:10124
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr466
		}
		goto st0
tr466:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st456
	st456:
		if p++; p == pe {
			goto _test_eof456
		}
	st_case_456:
//line rfc5424/parser.go:10143
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr467
		}
		goto st0
tr467:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st457
	st457:
		if p++; p == pe {
			goto _test_eof457
		}
	st_case_457:
//line rfc5424/parser.go:10162
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr468
		}
		goto st0
tr468:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st458
	st458:
		if p++; p == pe {
			goto _test_eof458
		}
	st_case_458:
//line rfc5424/parser.go:10181
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr469
		}
		goto st0
tr469:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st459
	st459:
		if p++; p == pe {
			goto _test_eof459
		}
	st_case_459:
//line rfc5424/parser.go:10200
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr470
		}
		goto st0
tr470:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st460
	st460:
		if p++; p == pe {
			goto _test_eof460
		}
	st_case_460:
//line rfc5424/parser.go:10219
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr471
		}
		goto st0
tr471:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st461
	st461:
		if p++; p == pe {
			goto _test_eof461
		}
	st_case_461:
//line rfc5424/parser.go:10238
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr472
		}
		goto st0
tr472:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st462
	st462:
		if p++; p == pe {
			goto _test_eof462
		}
	st_case_462:
//line rfc5424/parser.go:10257
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr473
		}
		goto st0
tr473:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st463
	st463:
		if p++; p == pe {
			goto _test_eof463
		}
	st_case_463:
//line rfc5424/parser.go:10276
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr474
		}
		goto st0
tr474:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st464
	st464:
		if p++; p == pe {
			goto _test_eof464
		}
	st_case_464:
//line rfc5424/parser.go:10295
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr475
		}
		goto st0
tr475:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st465
	st465:
		if p++; p == pe {
			goto _test_eof465
		}
	st_case_465:
//line rfc5424/parser.go:10314
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr476
		}
		goto st0
tr476:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st466
	st466:
		if p++; p == pe {
			goto _test_eof466
		}
	st_case_466:
//line rfc5424/parser.go:10333
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr477
		}
		goto st0
tr477:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st467
	st467:
		if p++; p == pe {
			goto _test_eof467
		}
	st_case_467:
//line rfc5424/parser.go:10352
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr478
		}
		goto st0
tr478:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st468
	st468:
		if p++; p == pe {
			goto _test_eof468
		}
	st_case_468:
//line rfc5424/parser.go:10371
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr479
		}
		goto st0
tr479:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st469
	st469:
		if p++; p == pe {
			goto _test_eof469
		}
	st_case_469:
//line rfc5424/parser.go:10390
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr480
		}
		goto st0
tr480:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st470
	st470:
		if p++; p == pe {
			goto _test_eof470
		}
	st_case_470:
//line rfc5424/parser.go:10409
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr481
		}
		goto st0
tr481:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st471
	st471:
		if p++; p == pe {
			goto _test_eof471
		}
	st_case_471:
//line rfc5424/parser.go:10428
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr482
		}
		goto st0
tr482:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st472
	st472:
		if p++; p == pe {
			goto _test_eof472
		}
	st_case_472:
//line rfc5424/parser.go:10447
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr483
		}
		goto st0
tr483:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st473
	st473:
		if p++; p == pe {
			goto _test_eof473
		}
	st_case_473:
//line rfc5424/parser.go:10466
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr484
		}
		goto st0
tr484:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st474
	st474:
		if p++; p == pe {
			goto _test_eof474
		}
	st_case_474:
//line rfc5424/parser.go:10485
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr485
		}
		goto st0
tr485:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st475
	st475:
		if p++; p == pe {
			goto _test_eof475
		}
	st_case_475:
//line rfc5424/parser.go:10504
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr486
		}
		goto st0
tr486:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st476
	st476:
		if p++; p == pe {
			goto _test_eof476
		}
	st_case_476:
//line rfc5424/parser.go:10523
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr487
		}
		goto st0
tr487:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st477
	st477:
		if p++; p == pe {
			goto _test_eof477
		}
	st_case_477:
//line rfc5424/parser.go:10542
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr488
		}
		goto st0
tr488:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st478
	st478:
		if p++; p == pe {
			goto _test_eof478
		}
	st_case_478:
//line rfc5424/parser.go:10561
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr489
		}
		goto st0
tr489:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st479
	st479:
		if p++; p == pe {
			goto _test_eof479
		}
	st_case_479:
//line rfc5424/parser.go:10580
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr490
		}
		goto st0
tr490:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st480
	st480:
		if p++; p == pe {
			goto _test_eof480
		}
	st_case_480:
//line rfc5424/parser.go:10599
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr491
		}
		goto st0
tr491:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st481
	st481:
		if p++; p == pe {
			goto _test_eof481
		}
	st_case_481:
//line rfc5424/parser.go:10618
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr492
		}
		goto st0
tr492:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st482
	st482:
		if p++; p == pe {
			goto _test_eof482
		}
	st_case_482:
//line rfc5424/parser.go:10637
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr493
		}
		goto st0
tr493:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st483
	st483:
		if p++; p == pe {
			goto _test_eof483
		}
	st_case_483:
//line rfc5424/parser.go:10656
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr494
		}
		goto st0
tr494:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st484
	st484:
		if p++; p == pe {
			goto _test_eof484
		}
	st_case_484:
//line rfc5424/parser.go:10675
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr495
		}
		goto st0
tr495:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st485
	st485:
		if p++; p == pe {
			goto _test_eof485
		}
	st_case_485:
//line rfc5424/parser.go:10694
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr496
		}
		goto st0
tr496:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st486
	st486:
		if p++; p == pe {
			goto _test_eof486
		}
	st_case_486:
//line rfc5424/parser.go:10713
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr497
		}
		goto st0
tr497:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st487
	st487:
		if p++; p == pe {
			goto _test_eof487
		}
	st_case_487:
//line rfc5424/parser.go:10732
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr498
		}
		goto st0
tr498:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st488
	st488:
		if p++; p == pe {
			goto _test_eof488
		}
	st_case_488:
//line rfc5424/parser.go:10751
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr499
		}
		goto st0
tr499:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st489
	st489:
		if p++; p == pe {
			goto _test_eof489
		}
	st_case_489:
//line rfc5424/parser.go:10770
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr500
		}
		goto st0
tr500:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st490
	st490:
		if p++; p == pe {
			goto _test_eof490
		}
	st_case_490:
//line rfc5424/parser.go:10789
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr501
		}
		goto st0
tr501:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st491
	st491:
		if p++; p == pe {
			goto _test_eof491
		}
	st_case_491:
//line rfc5424/parser.go:10808
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr502
		}
		goto st0
tr502:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st492
	st492:
		if p++; p == pe {
			goto _test_eof492
		}
	st_case_492:
//line rfc5424/parser.go:10827
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr503
		}
		goto st0
tr503:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st493
	st493:
		if p++; p == pe {
			goto _test_eof493
		}
	st_case_493:
//line rfc5424/parser.go:10846
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr504
		}
		goto st0
tr504:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st494
	st494:
		if p++; p == pe {
			goto _test_eof494
		}
	st_case_494:
//line rfc5424/parser.go:10865
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr505
		}
		goto st0
tr505:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st495
	st495:
		if p++; p == pe {
			goto _test_eof495
		}
	st_case_495:
//line rfc5424/parser.go:10884
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr506
		}
		goto st0
tr506:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st496
	st496:
		if p++; p == pe {
			goto _test_eof496
		}
	st_case_496:
//line rfc5424/parser.go:10903
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr507
		}
		goto st0
tr507:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st497
	st497:
		if p++; p == pe {
			goto _test_eof497
		}
	st_case_497:
//line rfc5424/parser.go:10922
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr508
		}
		goto st0
tr508:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st498
	st498:
		if p++; p == pe {
			goto _test_eof498
		}
	st_case_498:
//line rfc5424/parser.go:10941
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr509
		}
		goto st0
tr509:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st499
	st499:
		if p++; p == pe {
			goto _test_eof499
		}
	st_case_499:
//line rfc5424/parser.go:10960
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr510
		}
		goto st0
tr510:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st500
	st500:
		if p++; p == pe {
			goto _test_eof500
		}
	st_case_500:
//line rfc5424/parser.go:10979
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr511
		}
		goto st0
tr511:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st501
	st501:
		if p++; p == pe {
			goto _test_eof501
		}
	st_case_501:
//line rfc5424/parser.go:10998
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr512
		}
		goto st0
tr512:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st502
	st502:
		if p++; p == pe {
			goto _test_eof502
		}
	st_case_502:
//line rfc5424/parser.go:11017
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr513
		}
		goto st0
tr513:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st503
	st503:
		if p++; p == pe {
			goto _test_eof503
		}
	st_case_503:
//line rfc5424/parser.go:11036
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr514
		}
		goto st0
tr514:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st504
	st504:
		if p++; p == pe {
			goto _test_eof504
		}
	st_case_504:
//line rfc5424/parser.go:11055
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr515
		}
		goto st0
tr515:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st505
	st505:
		if p++; p == pe {
			goto _test_eof505
		}
	st_case_505:
//line rfc5424/parser.go:11074
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr516
		}
		goto st0
tr516:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st506
	st506:
		if p++; p == pe {
			goto _test_eof506
		}
	st_case_506:
//line rfc5424/parser.go:11093
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr517
		}
		goto st0
tr517:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st507
	st507:
		if p++; p == pe {
			goto _test_eof507
		}
	st_case_507:
//line rfc5424/parser.go:11112
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr518
		}
		goto st0
tr518:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st508
	st508:
		if p++; p == pe {
			goto _test_eof508
		}
	st_case_508:
//line rfc5424/parser.go:11131
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr519
		}
		goto st0
tr519:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st509
	st509:
		if p++; p == pe {
			goto _test_eof509
		}
	st_case_509:
//line rfc5424/parser.go:11150
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr520
		}
		goto st0
tr520:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st510
	st510:
		if p++; p == pe {
			goto _test_eof510
		}
	st_case_510:
//line rfc5424/parser.go:11169
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr521
		}
		goto st0
tr521:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st511
	st511:
		if p++; p == pe {
			goto _test_eof511
		}
	st_case_511:
//line rfc5424/parser.go:11188
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr522
		}
		goto st0
tr522:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st512
	st512:
		if p++; p == pe {
			goto _test_eof512
		}
	st_case_512:
//line rfc5424/parser.go:11207
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr523
		}
		goto st0
tr523:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st513
	st513:
		if p++; p == pe {
			goto _test_eof513
		}
	st_case_513:
//line rfc5424/parser.go:11226
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr524
		}
		goto st0
tr524:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st514
	st514:
		if p++; p == pe {
			goto _test_eof514
		}
	st_case_514:
//line rfc5424/parser.go:11245
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr525
		}
		goto st0
tr525:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st515
	st515:
		if p++; p == pe {
			goto _test_eof515
		}
	st_case_515:
//line rfc5424/parser.go:11264
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr526
		}
		goto st0
tr526:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st516
	st516:
		if p++; p == pe {
			goto _test_eof516
		}
	st_case_516:
//line rfc5424/parser.go:11283
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr527
		}
		goto st0
tr527:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st517
	st517:
		if p++; p == pe {
			goto _test_eof517
		}
	st_case_517:
//line rfc5424/parser.go:11302
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr528
		}
		goto st0
tr528:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st518
	st518:
		if p++; p == pe {
			goto _test_eof518
		}
	st_case_518:
//line rfc5424/parser.go:11321
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr529
		}
		goto st0
tr529:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st519
	st519:
		if p++; p == pe {
			goto _test_eof519
		}
	st_case_519:
//line rfc5424/parser.go:11340
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr530
		}
		goto st0
tr530:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st520
	st520:
		if p++; p == pe {
			goto _test_eof520
		}
	st_case_520:
//line rfc5424/parser.go:11359
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr531
		}
		goto st0
tr531:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st521
	st521:
		if p++; p == pe {
			goto _test_eof521
		}
	st_case_521:
//line rfc5424/parser.go:11378
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr532
		}
		goto st0
tr532:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st522
	st522:
		if p++; p == pe {
			goto _test_eof522
		}
	st_case_522:
//line rfc5424/parser.go:11397
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr533
		}
		goto st0
tr533:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st523
	st523:
		if p++; p == pe {
			goto _test_eof523
		}
	st_case_523:
//line rfc5424/parser.go:11416
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr534
		}
		goto st0
tr534:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st524
	st524:
		if p++; p == pe {
			goto _test_eof524
		}
	st_case_524:
//line rfc5424/parser.go:11435
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr535
		}
		goto st0
tr535:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st525
	st525:
		if p++; p == pe {
			goto _test_eof525
		}
	st_case_525:
//line rfc5424/parser.go:11454
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr536
		}
		goto st0
tr536:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st526
	st526:
		if p++; p == pe {
			goto _test_eof526
		}
	st_case_526:
//line rfc5424/parser.go:11473
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr537
		}
		goto st0
tr537:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st527
	st527:
		if p++; p == pe {
			goto _test_eof527
		}
	st_case_527:
//line rfc5424/parser.go:11492
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr538
		}
		goto st0
tr538:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st528
	st528:
		if p++; p == pe {
			goto _test_eof528
		}
	st_case_528:
//line rfc5424/parser.go:11511
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr539
		}
		goto st0
tr539:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st529
	st529:
		if p++; p == pe {
			goto _test_eof529
		}
	st_case_529:
//line rfc5424/parser.go:11530
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr540
		}
		goto st0
tr540:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st530
	st530:
		if p++; p == pe {
			goto _test_eof530
		}
	st_case_530:
//line rfc5424/parser.go:11549
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr541
		}
		goto st0
tr541:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st531
	st531:
		if p++; p == pe {
			goto _test_eof531
		}
	st_case_531:
//line rfc5424/parser.go:11568
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr542
		}
		goto st0
tr542:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st532
	st532:
		if p++; p == pe {
			goto _test_eof532
		}
	st_case_532:
//line rfc5424/parser.go:11587
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr543
		}
		goto st0
tr543:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st533
	st533:
		if p++; p == pe {
			goto _test_eof533
		}
	st_case_533:
//line rfc5424/parser.go:11606
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr544
		}
		goto st0
tr544:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st534
	st534:
		if p++; p == pe {
			goto _test_eof534
		}
	st_case_534:
//line rfc5424/parser.go:11625
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr545
		}
		goto st0
tr545:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st535
	st535:
		if p++; p == pe {
			goto _test_eof535
		}
	st_case_535:
//line rfc5424/parser.go:11644
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr546
		}
		goto st0
tr546:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st536
	st536:
		if p++; p == pe {
			goto _test_eof536
		}
	st_case_536:
//line rfc5424/parser.go:11663
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr547
		}
		goto st0
tr547:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st537
	st537:
		if p++; p == pe {
			goto _test_eof537
		}
	st_case_537:
//line rfc5424/parser.go:11682
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr548
		}
		goto st0
tr548:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st538
	st538:
		if p++; p == pe {
			goto _test_eof538
		}
	st_case_538:
//line rfc5424/parser.go:11701
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr549
		}
		goto st0
tr549:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st539
	st539:
		if p++; p == pe {
			goto _test_eof539
		}
	st_case_539:
//line rfc5424/parser.go:11720
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr550
		}
		goto st0
tr550:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st540
	st540:
		if p++; p == pe {
			goto _test_eof540
		}
	st_case_540:
//line rfc5424/parser.go:11739
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr551
		}
		goto st0
tr551:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st541
	st541:
		if p++; p == pe {
			goto _test_eof541
		}
	st_case_541:
//line rfc5424/parser.go:11758
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr552
		}
		goto st0
tr552:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st542
	st542:
		if p++; p == pe {
			goto _test_eof542
		}
	st_case_542:
//line rfc5424/parser.go:11777
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr553
		}
		goto st0
tr553:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st543
	st543:
		if p++; p == pe {
			goto _test_eof543
		}
	st_case_543:
//line rfc5424/parser.go:11796
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr554
		}
		goto st0
tr554:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st544
	st544:
		if p++; p == pe {
			goto _test_eof544
		}
	st_case_544:
//line rfc5424/parser.go:11815
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr555
		}
		goto st0
tr555:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st545
	st545:
		if p++; p == pe {
			goto _test_eof545
		}
	st_case_545:
//line rfc5424/parser.go:11834
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr556
		}
		goto st0
tr556:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st546
	st546:
		if p++; p == pe {
			goto _test_eof546
		}
	st_case_546:
//line rfc5424/parser.go:11853
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr557
		}
		goto st0
tr557:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st547
	st547:
		if p++; p == pe {
			goto _test_eof547
		}
	st_case_547:
//line rfc5424/parser.go:11872
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr558
		}
		goto st0
tr558:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st548
	st548:
		if p++; p == pe {
			goto _test_eof548
		}
	st_case_548:
//line rfc5424/parser.go:11891
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr559
		}
		goto st0
tr559:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st549
	st549:
		if p++; p == pe {
			goto _test_eof549
		}
	st_case_549:
//line rfc5424/parser.go:11910
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr560
		}
		goto st0
tr560:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st550
	st550:
		if p++; p == pe {
			goto _test_eof550
		}
	st_case_550:
//line rfc5424/parser.go:11929
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr561
		}
		goto st0
tr561:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st551
	st551:
		if p++; p == pe {
			goto _test_eof551
		}
	st_case_551:
//line rfc5424/parser.go:11948
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr562
		}
		goto st0
tr562:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st552
	st552:
		if p++; p == pe {
			goto _test_eof552
		}
	st_case_552:
//line rfc5424/parser.go:11967
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr563
		}
		goto st0
tr563:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st553
	st553:
		if p++; p == pe {
			goto _test_eof553
		}
	st_case_553:
//line rfc5424/parser.go:11986
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr564
		}
		goto st0
tr564:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st554
	st554:
		if p++; p == pe {
			goto _test_eof554
		}
	st_case_554:
//line rfc5424/parser.go:12005
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr565
		}
		goto st0
tr565:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st555
	st555:
		if p++; p == pe {
			goto _test_eof555
		}
	st_case_555:
//line rfc5424/parser.go:12024
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr566
		}
		goto st0
tr566:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st556
	st556:
		if p++; p == pe {
			goto _test_eof556
		}
	st_case_556:
//line rfc5424/parser.go:12043
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr567
		}
		goto st0
tr567:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st557
	st557:
		if p++; p == pe {
			goto _test_eof557
		}
	st_case_557:
//line rfc5424/parser.go:12062
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr568
		}
		goto st0
tr568:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st558
	st558:
		if p++; p == pe {
			goto _test_eof558
		}
	st_case_558:
//line rfc5424/parser.go:12081
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr569
		}
		goto st0
tr569:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st559
	st559:
		if p++; p == pe {
			goto _test_eof559
		}
	st_case_559:
//line rfc5424/parser.go:12100
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr570
		}
		goto st0
tr570:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st560
	st560:
		if p++; p == pe {
			goto _test_eof560
		}
	st_case_560:
//line rfc5424/parser.go:12119
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr571
		}
		goto st0
tr571:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st561
	st561:
		if p++; p == pe {
			goto _test_eof561
		}
	st_case_561:
//line rfc5424/parser.go:12138
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr572
		}
		goto st0
tr572:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st562
	st562:
		if p++; p == pe {
			goto _test_eof562
		}
	st_case_562:
//line rfc5424/parser.go:12157
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr573
		}
		goto st0
tr573:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st563
	st563:
		if p++; p == pe {
			goto _test_eof563
		}
	st_case_563:
//line rfc5424/parser.go:12176
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr574
		}
		goto st0
tr574:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st564
	st564:
		if p++; p == pe {
			goto _test_eof564
		}
	st_case_564:
//line rfc5424/parser.go:12195
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr575
		}
		goto st0
tr575:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st565
	st565:
		if p++; p == pe {
			goto _test_eof565
		}
	st_case_565:
//line rfc5424/parser.go:12214
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr576
		}
		goto st0
tr576:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st566
	st566:
		if p++; p == pe {
			goto _test_eof566
		}
	st_case_566:
//line rfc5424/parser.go:12233
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr577
		}
		goto st0
tr577:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st567
	st567:
		if p++; p == pe {
			goto _test_eof567
		}
	st_case_567:
//line rfc5424/parser.go:12252
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr578
		}
		goto st0
tr578:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st568
	st568:
		if p++; p == pe {
			goto _test_eof568
		}
	st_case_568:
//line rfc5424/parser.go:12271
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr579
		}
		goto st0
tr579:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st569
	st569:
		if p++; p == pe {
			goto _test_eof569
		}
	st_case_569:
//line rfc5424/parser.go:12290
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr580
		}
		goto st0
tr580:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st570
	st570:
		if p++; p == pe {
			goto _test_eof570
		}
	st_case_570:
//line rfc5424/parser.go:12309
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr581
		}
		goto st0
tr581:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st571
	st571:
		if p++; p == pe {
			goto _test_eof571
		}
	st_case_571:
//line rfc5424/parser.go:12328
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr582
		}
		goto st0
tr582:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st572
	st572:
		if p++; p == pe {
			goto _test_eof572
		}
	st_case_572:
//line rfc5424/parser.go:12347
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr583
		}
		goto st0
tr583:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st573
	st573:
		if p++; p == pe {
			goto _test_eof573
		}
	st_case_573:
//line rfc5424/parser.go:12366
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr584
		}
		goto st0
tr584:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st574
	st574:
		if p++; p == pe {
			goto _test_eof574
		}
	st_case_574:
//line rfc5424/parser.go:12385
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr585
		}
		goto st0
tr585:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st575
	st575:
		if p++; p == pe {
			goto _test_eof575
		}
	st_case_575:
//line rfc5424/parser.go:12404
		if data[p] == 32 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr586
		}
		goto st0
tr586:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st576
	st576:
		if p++; p == pe {
			goto _test_eof576
		}
	st_case_576:
//line rfc5424/parser.go:12423
		if data[p] == 32 {
			goto tr15
		}
		goto st0
tr11:
//line rfc5424/machine.rl:49

    err = fmt.Errorf("error parsing <nilvalue>");

//line rfc5424/machine.rl:36

    poss["timestamp:ini"] = p

	goto st577
	st577:
		if p++; p == pe {
			goto _test_eof577
		}
	st_case_577:
//line rfc5424/parser.go:12443
		if 48 <= data[p] && data[p] <= 57 {
			goto st578
		}
		goto st0
	st578:
		if p++; p == pe {
			goto _test_eof578
		}
	st_case_578:
		if 48 <= data[p] && data[p] <= 57 {
			goto st579
		}
		goto st0
	st579:
		if p++; p == pe {
			goto _test_eof579
		}
	st_case_579:
		if 48 <= data[p] && data[p] <= 57 {
			goto st580
		}
		goto st0
	st580:
		if p++; p == pe {
			goto _test_eof580
		}
	st_case_580:
		if data[p] == 45 {
			goto st581
		}
		goto st0
	st581:
		if p++; p == pe {
			goto _test_eof581
		}
	st_case_581:
		switch data[p] {
		case 48:
			goto st582
		case 49:
			goto st613
		}
		goto st0
	st582:
		if p++; p == pe {
			goto _test_eof582
		}
	st_case_582:
		if 49 <= data[p] && data[p] <= 57 {
			goto st583
		}
		goto st0
	st583:
		if p++; p == pe {
			goto _test_eof583
		}
	st_case_583:
		if data[p] == 45 {
			goto st584
		}
		goto st0
	st584:
		if p++; p == pe {
			goto _test_eof584
		}
	st_case_584:
		switch data[p] {
		case 48:
			goto st585
		case 51:
			goto st612
		}
		if 49 <= data[p] && data[p] <= 50 {
			goto st611
		}
		goto st0
	st585:
		if p++; p == pe {
			goto _test_eof585
		}
	st_case_585:
		if 49 <= data[p] && data[p] <= 57 {
			goto st586
		}
		goto st0
	st586:
		if p++; p == pe {
			goto _test_eof586
		}
	st_case_586:
		if data[p] == 84 {
			goto st587
		}
		goto st0
	st587:
		if p++; p == pe {
			goto _test_eof587
		}
	st_case_587:
		if data[p] == 50 {
			goto st610
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st588
		}
		goto st0
	st588:
		if p++; p == pe {
			goto _test_eof588
		}
	st_case_588:
		if 48 <= data[p] && data[p] <= 57 {
			goto st589
		}
		goto st0
	st589:
		if p++; p == pe {
			goto _test_eof589
		}
	st_case_589:
		if data[p] == 58 {
			goto st590
		}
		goto st0
	st590:
		if p++; p == pe {
			goto _test_eof590
		}
	st_case_590:
		if 48 <= data[p] && data[p] <= 53 {
			goto st591
		}
		goto st0
	st591:
		if p++; p == pe {
			goto _test_eof591
		}
	st_case_591:
		if 48 <= data[p] && data[p] <= 57 {
			goto st592
		}
		goto st0
	st592:
		if p++; p == pe {
			goto _test_eof592
		}
	st_case_592:
		if data[p] == 58 {
			goto st593
		}
		goto st0
	st593:
		if p++; p == pe {
			goto _test_eof593
		}
	st_case_593:
		if 48 <= data[p] && data[p] <= 53 {
			goto st594
		}
		goto st0
	st594:
		if p++; p == pe {
			goto _test_eof594
		}
	st_case_594:
		if 48 <= data[p] && data[p] <= 57 {
			goto st595
		}
		goto st0
	st595:
		if p++; p == pe {
			goto _test_eof595
		}
	st_case_595:
		switch data[p] {
		case 43:
			goto st596
		case 45:
			goto st596
		case 46:
			goto st603
		case 90:
			goto st601
		}
		goto st0
	st596:
		if p++; p == pe {
			goto _test_eof596
		}
	st_case_596:
		if data[p] == 50 {
			goto st602
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st597
		}
		goto st0
	st597:
		if p++; p == pe {
			goto _test_eof597
		}
	st_case_597:
		if 48 <= data[p] && data[p] <= 57 {
			goto st598
		}
		goto st0
	st598:
		if p++; p == pe {
			goto _test_eof598
		}
	st_case_598:
		if data[p] == 58 {
			goto st599
		}
		goto st0
	st599:
		if p++; p == pe {
			goto _test_eof599
		}
	st_case_599:
		if 48 <= data[p] && data[p] <= 53 {
			goto st600
		}
		goto st0
	st600:
		if p++; p == pe {
			goto _test_eof600
		}
	st_case_600:
		if 48 <= data[p] && data[p] <= 57 {
			goto st601
		}
		goto st0
	st601:
		if p++; p == pe {
			goto _test_eof601
		}
	st_case_601:
		if data[p] == 32 {
			goto tr617
		}
		goto st0
	st602:
		if p++; p == pe {
			goto _test_eof602
		}
	st_case_602:
		if 48 <= data[p] && data[p] <= 51 {
			goto st598
		}
		goto st0
	st603:
		if p++; p == pe {
			goto _test_eof603
		}
	st_case_603:
		if 48 <= data[p] && data[p] <= 57 {
			goto st604
		}
		goto st0
	st604:
		if p++; p == pe {
			goto _test_eof604
		}
	st_case_604:
		switch data[p] {
		case 43:
			goto st596
		case 45:
			goto st596
		case 90:
			goto st601
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st605
		}
		goto st0
	st605:
		if p++; p == pe {
			goto _test_eof605
		}
	st_case_605:
		switch data[p] {
		case 43:
			goto st596
		case 45:
			goto st596
		case 90:
			goto st601
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st606
		}
		goto st0
	st606:
		if p++; p == pe {
			goto _test_eof606
		}
	st_case_606:
		switch data[p] {
		case 43:
			goto st596
		case 45:
			goto st596
		case 90:
			goto st601
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st607
		}
		goto st0
	st607:
		if p++; p == pe {
			goto _test_eof607
		}
	st_case_607:
		switch data[p] {
		case 43:
			goto st596
		case 45:
			goto st596
		case 90:
			goto st601
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st608
		}
		goto st0
	st608:
		if p++; p == pe {
			goto _test_eof608
		}
	st_case_608:
		switch data[p] {
		case 43:
			goto st596
		case 45:
			goto st596
		case 90:
			goto st601
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st609
		}
		goto st0
	st609:
		if p++; p == pe {
			goto _test_eof609
		}
	st_case_609:
		switch data[p] {
		case 43:
			goto st596
		case 45:
			goto st596
		case 90:
			goto st601
		}
		goto st0
	st610:
		if p++; p == pe {
			goto _test_eof610
		}
	st_case_610:
		if 48 <= data[p] && data[p] <= 51 {
			goto st589
		}
		goto st0
	st611:
		if p++; p == pe {
			goto _test_eof611
		}
	st_case_611:
		if 48 <= data[p] && data[p] <= 57 {
			goto st586
		}
		goto st0
	st612:
		if p++; p == pe {
			goto _test_eof612
		}
	st_case_612:
		if 48 <= data[p] && data[p] <= 49 {
			goto st586
		}
		goto st0
	st613:
		if p++; p == pe {
			goto _test_eof613
		}
	st_case_613:
		if 48 <= data[p] && data[p] <= 50 {
			goto st583
		}
		goto st0
tr8:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st614
	st614:
		if p++; p == pe {
			goto _test_eof614
		}
	st_case_614:
//line rfc5424/parser.go:12850
		if data[p] == 32 {
			goto tr7
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr624
		}
		goto st0
tr624:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st615
	st615:
		if p++; p == pe {
			goto _test_eof615
		}
	st_case_615:
//line rfc5424/parser.go:12869
		if data[p] == 32 {
			goto tr7
		}
		goto st0
tr3:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st616
	st616:
		if p++; p == pe {
			goto _test_eof616
		}
	st_case_616:
//line rfc5424/parser.go:12885
		switch data[p] {
		case 57:
			goto tr625
		case 62:
			goto tr5
		}
		if 48 <= data[p] && data[p] <= 56 {
			goto tr4
		}
		goto st0
tr4:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st617
	st617:
		if p++; p == pe {
			goto _test_eof617
		}
	st_case_617:
//line rfc5424/parser.go:12907
		if data[p] == 62 {
			goto tr5
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr2
		}
		goto st0
tr625:
//line rfc5424/machine.rl:7

    cr.Add(data[p])

	goto st618
	st618:
		if p++; p == pe {
			goto _test_eof618
		}
	st_case_618:
//line rfc5424/parser.go:12926
		if data[p] == 62 {
			goto tr5
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto tr2
		}
		goto st0
tr629:
//line rfc5424/machine.rl:127
 {goto st1 } 
	goto st623
	st623:
		if p++; p == pe {
			goto _test_eof623
		}
	st_case_623:
//line rfc5424/parser.go:12943
		switch data[p] {
		case 10:
			goto st0
		case 13:
			goto st0
		}
		goto tr629
	st_out:
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof619: cs = 619; goto _test_eof
	_test_eof620: cs = 620; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof621: cs = 621; goto _test_eof
	_test_eof55: cs = 55; goto _test_eof
	_test_eof56: cs = 56; goto _test_eof
	_test_eof57: cs = 57; goto _test_eof
	_test_eof58: cs = 58; goto _test_eof
	_test_eof59: cs = 59; goto _test_eof
	_test_eof60: cs = 60; goto _test_eof
	_test_eof61: cs = 61; goto _test_eof
	_test_eof62: cs = 62; goto _test_eof
	_test_eof63: cs = 63; goto _test_eof
	_test_eof64: cs = 64; goto _test_eof
	_test_eof65: cs = 65; goto _test_eof
	_test_eof66: cs = 66; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof70: cs = 70; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof81: cs = 81; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof83: cs = 83; goto _test_eof
	_test_eof84: cs = 84; goto _test_eof
	_test_eof85: cs = 85; goto _test_eof
	_test_eof86: cs = 86; goto _test_eof
	_test_eof87: cs = 87; goto _test_eof
	_test_eof88: cs = 88; goto _test_eof
	_test_eof89: cs = 89; goto _test_eof
	_test_eof90: cs = 90; goto _test_eof
	_test_eof91: cs = 91; goto _test_eof
	_test_eof92: cs = 92; goto _test_eof
	_test_eof93: cs = 93; goto _test_eof
	_test_eof94: cs = 94; goto _test_eof
	_test_eof95: cs = 95; goto _test_eof
	_test_eof96: cs = 96; goto _test_eof
	_test_eof97: cs = 97; goto _test_eof
	_test_eof98: cs = 98; goto _test_eof
	_test_eof99: cs = 99; goto _test_eof
	_test_eof100: cs = 100; goto _test_eof
	_test_eof101: cs = 101; goto _test_eof
	_test_eof102: cs = 102; goto _test_eof
	_test_eof103: cs = 103; goto _test_eof
	_test_eof104: cs = 104; goto _test_eof
	_test_eof105: cs = 105; goto _test_eof
	_test_eof106: cs = 106; goto _test_eof
	_test_eof107: cs = 107; goto _test_eof
	_test_eof108: cs = 108; goto _test_eof
	_test_eof109: cs = 109; goto _test_eof
	_test_eof110: cs = 110; goto _test_eof
	_test_eof111: cs = 111; goto _test_eof
	_test_eof112: cs = 112; goto _test_eof
	_test_eof113: cs = 113; goto _test_eof
	_test_eof114: cs = 114; goto _test_eof
	_test_eof115: cs = 115; goto _test_eof
	_test_eof116: cs = 116; goto _test_eof
	_test_eof117: cs = 117; goto _test_eof
	_test_eof622: cs = 622; goto _test_eof
	_test_eof118: cs = 118; goto _test_eof
	_test_eof119: cs = 119; goto _test_eof
	_test_eof120: cs = 120; goto _test_eof
	_test_eof121: cs = 121; goto _test_eof
	_test_eof122: cs = 122; goto _test_eof
	_test_eof123: cs = 123; goto _test_eof
	_test_eof124: cs = 124; goto _test_eof
	_test_eof125: cs = 125; goto _test_eof
	_test_eof126: cs = 126; goto _test_eof
	_test_eof127: cs = 127; goto _test_eof
	_test_eof128: cs = 128; goto _test_eof
	_test_eof129: cs = 129; goto _test_eof
	_test_eof130: cs = 130; goto _test_eof
	_test_eof131: cs = 131; goto _test_eof
	_test_eof132: cs = 132; goto _test_eof
	_test_eof133: cs = 133; goto _test_eof
	_test_eof134: cs = 134; goto _test_eof
	_test_eof135: cs = 135; goto _test_eof
	_test_eof136: cs = 136; goto _test_eof
	_test_eof137: cs = 137; goto _test_eof
	_test_eof138: cs = 138; goto _test_eof
	_test_eof139: cs = 139; goto _test_eof
	_test_eof140: cs = 140; goto _test_eof
	_test_eof141: cs = 141; goto _test_eof
	_test_eof142: cs = 142; goto _test_eof
	_test_eof143: cs = 143; goto _test_eof
	_test_eof144: cs = 144; goto _test_eof
	_test_eof145: cs = 145; goto _test_eof
	_test_eof146: cs = 146; goto _test_eof
	_test_eof147: cs = 147; goto _test_eof
	_test_eof148: cs = 148; goto _test_eof
	_test_eof149: cs = 149; goto _test_eof
	_test_eof150: cs = 150; goto _test_eof
	_test_eof151: cs = 151; goto _test_eof
	_test_eof152: cs = 152; goto _test_eof
	_test_eof153: cs = 153; goto _test_eof
	_test_eof154: cs = 154; goto _test_eof
	_test_eof155: cs = 155; goto _test_eof
	_test_eof156: cs = 156; goto _test_eof
	_test_eof157: cs = 157; goto _test_eof
	_test_eof158: cs = 158; goto _test_eof
	_test_eof159: cs = 159; goto _test_eof
	_test_eof160: cs = 160; goto _test_eof
	_test_eof161: cs = 161; goto _test_eof
	_test_eof162: cs = 162; goto _test_eof
	_test_eof163: cs = 163; goto _test_eof
	_test_eof164: cs = 164; goto _test_eof
	_test_eof165: cs = 165; goto _test_eof
	_test_eof166: cs = 166; goto _test_eof
	_test_eof167: cs = 167; goto _test_eof
	_test_eof168: cs = 168; goto _test_eof
	_test_eof169: cs = 169; goto _test_eof
	_test_eof170: cs = 170; goto _test_eof
	_test_eof171: cs = 171; goto _test_eof
	_test_eof172: cs = 172; goto _test_eof
	_test_eof173: cs = 173; goto _test_eof
	_test_eof174: cs = 174; goto _test_eof
	_test_eof175: cs = 175; goto _test_eof
	_test_eof176: cs = 176; goto _test_eof
	_test_eof177: cs = 177; goto _test_eof
	_test_eof178: cs = 178; goto _test_eof
	_test_eof179: cs = 179; goto _test_eof
	_test_eof180: cs = 180; goto _test_eof
	_test_eof181: cs = 181; goto _test_eof
	_test_eof182: cs = 182; goto _test_eof
	_test_eof183: cs = 183; goto _test_eof
	_test_eof184: cs = 184; goto _test_eof
	_test_eof185: cs = 185; goto _test_eof
	_test_eof186: cs = 186; goto _test_eof
	_test_eof187: cs = 187; goto _test_eof
	_test_eof188: cs = 188; goto _test_eof
	_test_eof189: cs = 189; goto _test_eof
	_test_eof190: cs = 190; goto _test_eof
	_test_eof191: cs = 191; goto _test_eof
	_test_eof192: cs = 192; goto _test_eof
	_test_eof193: cs = 193; goto _test_eof
	_test_eof194: cs = 194; goto _test_eof
	_test_eof195: cs = 195; goto _test_eof
	_test_eof196: cs = 196; goto _test_eof
	_test_eof197: cs = 197; goto _test_eof
	_test_eof198: cs = 198; goto _test_eof
	_test_eof199: cs = 199; goto _test_eof
	_test_eof200: cs = 200; goto _test_eof
	_test_eof201: cs = 201; goto _test_eof
	_test_eof202: cs = 202; goto _test_eof
	_test_eof203: cs = 203; goto _test_eof
	_test_eof204: cs = 204; goto _test_eof
	_test_eof205: cs = 205; goto _test_eof
	_test_eof206: cs = 206; goto _test_eof
	_test_eof207: cs = 207; goto _test_eof
	_test_eof208: cs = 208; goto _test_eof
	_test_eof209: cs = 209; goto _test_eof
	_test_eof210: cs = 210; goto _test_eof
	_test_eof211: cs = 211; goto _test_eof
	_test_eof212: cs = 212; goto _test_eof
	_test_eof213: cs = 213; goto _test_eof
	_test_eof214: cs = 214; goto _test_eof
	_test_eof215: cs = 215; goto _test_eof
	_test_eof216: cs = 216; goto _test_eof
	_test_eof217: cs = 217; goto _test_eof
	_test_eof218: cs = 218; goto _test_eof
	_test_eof219: cs = 219; goto _test_eof
	_test_eof220: cs = 220; goto _test_eof
	_test_eof221: cs = 221; goto _test_eof
	_test_eof222: cs = 222; goto _test_eof
	_test_eof223: cs = 223; goto _test_eof
	_test_eof224: cs = 224; goto _test_eof
	_test_eof225: cs = 225; goto _test_eof
	_test_eof226: cs = 226; goto _test_eof
	_test_eof227: cs = 227; goto _test_eof
	_test_eof228: cs = 228; goto _test_eof
	_test_eof229: cs = 229; goto _test_eof
	_test_eof230: cs = 230; goto _test_eof
	_test_eof231: cs = 231; goto _test_eof
	_test_eof232: cs = 232; goto _test_eof
	_test_eof233: cs = 233; goto _test_eof
	_test_eof234: cs = 234; goto _test_eof
	_test_eof235: cs = 235; goto _test_eof
	_test_eof236: cs = 236; goto _test_eof
	_test_eof237: cs = 237; goto _test_eof
	_test_eof238: cs = 238; goto _test_eof
	_test_eof239: cs = 239; goto _test_eof
	_test_eof240: cs = 240; goto _test_eof
	_test_eof241: cs = 241; goto _test_eof
	_test_eof242: cs = 242; goto _test_eof
	_test_eof243: cs = 243; goto _test_eof
	_test_eof244: cs = 244; goto _test_eof
	_test_eof245: cs = 245; goto _test_eof
	_test_eof246: cs = 246; goto _test_eof
	_test_eof247: cs = 247; goto _test_eof
	_test_eof248: cs = 248; goto _test_eof
	_test_eof249: cs = 249; goto _test_eof
	_test_eof250: cs = 250; goto _test_eof
	_test_eof251: cs = 251; goto _test_eof
	_test_eof252: cs = 252; goto _test_eof
	_test_eof253: cs = 253; goto _test_eof
	_test_eof254: cs = 254; goto _test_eof
	_test_eof255: cs = 255; goto _test_eof
	_test_eof256: cs = 256; goto _test_eof
	_test_eof257: cs = 257; goto _test_eof
	_test_eof258: cs = 258; goto _test_eof
	_test_eof259: cs = 259; goto _test_eof
	_test_eof260: cs = 260; goto _test_eof
	_test_eof261: cs = 261; goto _test_eof
	_test_eof262: cs = 262; goto _test_eof
	_test_eof263: cs = 263; goto _test_eof
	_test_eof264: cs = 264; goto _test_eof
	_test_eof265: cs = 265; goto _test_eof
	_test_eof266: cs = 266; goto _test_eof
	_test_eof267: cs = 267; goto _test_eof
	_test_eof268: cs = 268; goto _test_eof
	_test_eof269: cs = 269; goto _test_eof
	_test_eof270: cs = 270; goto _test_eof
	_test_eof271: cs = 271; goto _test_eof
	_test_eof272: cs = 272; goto _test_eof
	_test_eof273: cs = 273; goto _test_eof
	_test_eof274: cs = 274; goto _test_eof
	_test_eof275: cs = 275; goto _test_eof
	_test_eof276: cs = 276; goto _test_eof
	_test_eof277: cs = 277; goto _test_eof
	_test_eof278: cs = 278; goto _test_eof
	_test_eof279: cs = 279; goto _test_eof
	_test_eof280: cs = 280; goto _test_eof
	_test_eof281: cs = 281; goto _test_eof
	_test_eof282: cs = 282; goto _test_eof
	_test_eof283: cs = 283; goto _test_eof
	_test_eof284: cs = 284; goto _test_eof
	_test_eof285: cs = 285; goto _test_eof
	_test_eof286: cs = 286; goto _test_eof
	_test_eof287: cs = 287; goto _test_eof
	_test_eof288: cs = 288; goto _test_eof
	_test_eof289: cs = 289; goto _test_eof
	_test_eof290: cs = 290; goto _test_eof
	_test_eof291: cs = 291; goto _test_eof
	_test_eof292: cs = 292; goto _test_eof
	_test_eof293: cs = 293; goto _test_eof
	_test_eof294: cs = 294; goto _test_eof
	_test_eof295: cs = 295; goto _test_eof
	_test_eof296: cs = 296; goto _test_eof
	_test_eof297: cs = 297; goto _test_eof
	_test_eof298: cs = 298; goto _test_eof
	_test_eof299: cs = 299; goto _test_eof
	_test_eof300: cs = 300; goto _test_eof
	_test_eof301: cs = 301; goto _test_eof
	_test_eof302: cs = 302; goto _test_eof
	_test_eof303: cs = 303; goto _test_eof
	_test_eof304: cs = 304; goto _test_eof
	_test_eof305: cs = 305; goto _test_eof
	_test_eof306: cs = 306; goto _test_eof
	_test_eof307: cs = 307; goto _test_eof
	_test_eof308: cs = 308; goto _test_eof
	_test_eof309: cs = 309; goto _test_eof
	_test_eof310: cs = 310; goto _test_eof
	_test_eof311: cs = 311; goto _test_eof
	_test_eof312: cs = 312; goto _test_eof
	_test_eof313: cs = 313; goto _test_eof
	_test_eof314: cs = 314; goto _test_eof
	_test_eof315: cs = 315; goto _test_eof
	_test_eof316: cs = 316; goto _test_eof
	_test_eof317: cs = 317; goto _test_eof
	_test_eof318: cs = 318; goto _test_eof
	_test_eof319: cs = 319; goto _test_eof
	_test_eof320: cs = 320; goto _test_eof
	_test_eof321: cs = 321; goto _test_eof
	_test_eof322: cs = 322; goto _test_eof
	_test_eof323: cs = 323; goto _test_eof
	_test_eof324: cs = 324; goto _test_eof
	_test_eof325: cs = 325; goto _test_eof
	_test_eof326: cs = 326; goto _test_eof
	_test_eof327: cs = 327; goto _test_eof
	_test_eof328: cs = 328; goto _test_eof
	_test_eof329: cs = 329; goto _test_eof
	_test_eof330: cs = 330; goto _test_eof
	_test_eof331: cs = 331; goto _test_eof
	_test_eof332: cs = 332; goto _test_eof
	_test_eof333: cs = 333; goto _test_eof
	_test_eof334: cs = 334; goto _test_eof
	_test_eof335: cs = 335; goto _test_eof
	_test_eof336: cs = 336; goto _test_eof
	_test_eof337: cs = 337; goto _test_eof
	_test_eof338: cs = 338; goto _test_eof
	_test_eof339: cs = 339; goto _test_eof
	_test_eof340: cs = 340; goto _test_eof
	_test_eof341: cs = 341; goto _test_eof
	_test_eof342: cs = 342; goto _test_eof
	_test_eof343: cs = 343; goto _test_eof
	_test_eof344: cs = 344; goto _test_eof
	_test_eof345: cs = 345; goto _test_eof
	_test_eof346: cs = 346; goto _test_eof
	_test_eof347: cs = 347; goto _test_eof
	_test_eof348: cs = 348; goto _test_eof
	_test_eof349: cs = 349; goto _test_eof
	_test_eof350: cs = 350; goto _test_eof
	_test_eof351: cs = 351; goto _test_eof
	_test_eof352: cs = 352; goto _test_eof
	_test_eof353: cs = 353; goto _test_eof
	_test_eof354: cs = 354; goto _test_eof
	_test_eof355: cs = 355; goto _test_eof
	_test_eof356: cs = 356; goto _test_eof
	_test_eof357: cs = 357; goto _test_eof
	_test_eof358: cs = 358; goto _test_eof
	_test_eof359: cs = 359; goto _test_eof
	_test_eof360: cs = 360; goto _test_eof
	_test_eof361: cs = 361; goto _test_eof
	_test_eof362: cs = 362; goto _test_eof
	_test_eof363: cs = 363; goto _test_eof
	_test_eof364: cs = 364; goto _test_eof
	_test_eof365: cs = 365; goto _test_eof
	_test_eof366: cs = 366; goto _test_eof
	_test_eof367: cs = 367; goto _test_eof
	_test_eof368: cs = 368; goto _test_eof
	_test_eof369: cs = 369; goto _test_eof
	_test_eof370: cs = 370; goto _test_eof
	_test_eof371: cs = 371; goto _test_eof
	_test_eof372: cs = 372; goto _test_eof
	_test_eof373: cs = 373; goto _test_eof
	_test_eof374: cs = 374; goto _test_eof
	_test_eof375: cs = 375; goto _test_eof
	_test_eof376: cs = 376; goto _test_eof
	_test_eof377: cs = 377; goto _test_eof
	_test_eof378: cs = 378; goto _test_eof
	_test_eof379: cs = 379; goto _test_eof
	_test_eof380: cs = 380; goto _test_eof
	_test_eof381: cs = 381; goto _test_eof
	_test_eof382: cs = 382; goto _test_eof
	_test_eof383: cs = 383; goto _test_eof
	_test_eof384: cs = 384; goto _test_eof
	_test_eof385: cs = 385; goto _test_eof
	_test_eof386: cs = 386; goto _test_eof
	_test_eof387: cs = 387; goto _test_eof
	_test_eof388: cs = 388; goto _test_eof
	_test_eof389: cs = 389; goto _test_eof
	_test_eof390: cs = 390; goto _test_eof
	_test_eof391: cs = 391; goto _test_eof
	_test_eof392: cs = 392; goto _test_eof
	_test_eof393: cs = 393; goto _test_eof
	_test_eof394: cs = 394; goto _test_eof
	_test_eof395: cs = 395; goto _test_eof
	_test_eof396: cs = 396; goto _test_eof
	_test_eof397: cs = 397; goto _test_eof
	_test_eof398: cs = 398; goto _test_eof
	_test_eof399: cs = 399; goto _test_eof
	_test_eof400: cs = 400; goto _test_eof
	_test_eof401: cs = 401; goto _test_eof
	_test_eof402: cs = 402; goto _test_eof
	_test_eof403: cs = 403; goto _test_eof
	_test_eof404: cs = 404; goto _test_eof
	_test_eof405: cs = 405; goto _test_eof
	_test_eof406: cs = 406; goto _test_eof
	_test_eof407: cs = 407; goto _test_eof
	_test_eof408: cs = 408; goto _test_eof
	_test_eof409: cs = 409; goto _test_eof
	_test_eof410: cs = 410; goto _test_eof
	_test_eof411: cs = 411; goto _test_eof
	_test_eof412: cs = 412; goto _test_eof
	_test_eof413: cs = 413; goto _test_eof
	_test_eof414: cs = 414; goto _test_eof
	_test_eof415: cs = 415; goto _test_eof
	_test_eof416: cs = 416; goto _test_eof
	_test_eof417: cs = 417; goto _test_eof
	_test_eof418: cs = 418; goto _test_eof
	_test_eof419: cs = 419; goto _test_eof
	_test_eof420: cs = 420; goto _test_eof
	_test_eof421: cs = 421; goto _test_eof
	_test_eof422: cs = 422; goto _test_eof
	_test_eof423: cs = 423; goto _test_eof
	_test_eof424: cs = 424; goto _test_eof
	_test_eof425: cs = 425; goto _test_eof
	_test_eof426: cs = 426; goto _test_eof
	_test_eof427: cs = 427; goto _test_eof
	_test_eof428: cs = 428; goto _test_eof
	_test_eof429: cs = 429; goto _test_eof
	_test_eof430: cs = 430; goto _test_eof
	_test_eof431: cs = 431; goto _test_eof
	_test_eof432: cs = 432; goto _test_eof
	_test_eof433: cs = 433; goto _test_eof
	_test_eof434: cs = 434; goto _test_eof
	_test_eof435: cs = 435; goto _test_eof
	_test_eof436: cs = 436; goto _test_eof
	_test_eof437: cs = 437; goto _test_eof
	_test_eof438: cs = 438; goto _test_eof
	_test_eof439: cs = 439; goto _test_eof
	_test_eof440: cs = 440; goto _test_eof
	_test_eof441: cs = 441; goto _test_eof
	_test_eof442: cs = 442; goto _test_eof
	_test_eof443: cs = 443; goto _test_eof
	_test_eof444: cs = 444; goto _test_eof
	_test_eof445: cs = 445; goto _test_eof
	_test_eof446: cs = 446; goto _test_eof
	_test_eof447: cs = 447; goto _test_eof
	_test_eof448: cs = 448; goto _test_eof
	_test_eof449: cs = 449; goto _test_eof
	_test_eof450: cs = 450; goto _test_eof
	_test_eof451: cs = 451; goto _test_eof
	_test_eof452: cs = 452; goto _test_eof
	_test_eof453: cs = 453; goto _test_eof
	_test_eof454: cs = 454; goto _test_eof
	_test_eof455: cs = 455; goto _test_eof
	_test_eof456: cs = 456; goto _test_eof
	_test_eof457: cs = 457; goto _test_eof
	_test_eof458: cs = 458; goto _test_eof
	_test_eof459: cs = 459; goto _test_eof
	_test_eof460: cs = 460; goto _test_eof
	_test_eof461: cs = 461; goto _test_eof
	_test_eof462: cs = 462; goto _test_eof
	_test_eof463: cs = 463; goto _test_eof
	_test_eof464: cs = 464; goto _test_eof
	_test_eof465: cs = 465; goto _test_eof
	_test_eof466: cs = 466; goto _test_eof
	_test_eof467: cs = 467; goto _test_eof
	_test_eof468: cs = 468; goto _test_eof
	_test_eof469: cs = 469; goto _test_eof
	_test_eof470: cs = 470; goto _test_eof
	_test_eof471: cs = 471; goto _test_eof
	_test_eof472: cs = 472; goto _test_eof
	_test_eof473: cs = 473; goto _test_eof
	_test_eof474: cs = 474; goto _test_eof
	_test_eof475: cs = 475; goto _test_eof
	_test_eof476: cs = 476; goto _test_eof
	_test_eof477: cs = 477; goto _test_eof
	_test_eof478: cs = 478; goto _test_eof
	_test_eof479: cs = 479; goto _test_eof
	_test_eof480: cs = 480; goto _test_eof
	_test_eof481: cs = 481; goto _test_eof
	_test_eof482: cs = 482; goto _test_eof
	_test_eof483: cs = 483; goto _test_eof
	_test_eof484: cs = 484; goto _test_eof
	_test_eof485: cs = 485; goto _test_eof
	_test_eof486: cs = 486; goto _test_eof
	_test_eof487: cs = 487; goto _test_eof
	_test_eof488: cs = 488; goto _test_eof
	_test_eof489: cs = 489; goto _test_eof
	_test_eof490: cs = 490; goto _test_eof
	_test_eof491: cs = 491; goto _test_eof
	_test_eof492: cs = 492; goto _test_eof
	_test_eof493: cs = 493; goto _test_eof
	_test_eof494: cs = 494; goto _test_eof
	_test_eof495: cs = 495; goto _test_eof
	_test_eof496: cs = 496; goto _test_eof
	_test_eof497: cs = 497; goto _test_eof
	_test_eof498: cs = 498; goto _test_eof
	_test_eof499: cs = 499; goto _test_eof
	_test_eof500: cs = 500; goto _test_eof
	_test_eof501: cs = 501; goto _test_eof
	_test_eof502: cs = 502; goto _test_eof
	_test_eof503: cs = 503; goto _test_eof
	_test_eof504: cs = 504; goto _test_eof
	_test_eof505: cs = 505; goto _test_eof
	_test_eof506: cs = 506; goto _test_eof
	_test_eof507: cs = 507; goto _test_eof
	_test_eof508: cs = 508; goto _test_eof
	_test_eof509: cs = 509; goto _test_eof
	_test_eof510: cs = 510; goto _test_eof
	_test_eof511: cs = 511; goto _test_eof
	_test_eof512: cs = 512; goto _test_eof
	_test_eof513: cs = 513; goto _test_eof
	_test_eof514: cs = 514; goto _test_eof
	_test_eof515: cs = 515; goto _test_eof
	_test_eof516: cs = 516; goto _test_eof
	_test_eof517: cs = 517; goto _test_eof
	_test_eof518: cs = 518; goto _test_eof
	_test_eof519: cs = 519; goto _test_eof
	_test_eof520: cs = 520; goto _test_eof
	_test_eof521: cs = 521; goto _test_eof
	_test_eof522: cs = 522; goto _test_eof
	_test_eof523: cs = 523; goto _test_eof
	_test_eof524: cs = 524; goto _test_eof
	_test_eof525: cs = 525; goto _test_eof
	_test_eof526: cs = 526; goto _test_eof
	_test_eof527: cs = 527; goto _test_eof
	_test_eof528: cs = 528; goto _test_eof
	_test_eof529: cs = 529; goto _test_eof
	_test_eof530: cs = 530; goto _test_eof
	_test_eof531: cs = 531; goto _test_eof
	_test_eof532: cs = 532; goto _test_eof
	_test_eof533: cs = 533; goto _test_eof
	_test_eof534: cs = 534; goto _test_eof
	_test_eof535: cs = 535; goto _test_eof
	_test_eof536: cs = 536; goto _test_eof
	_test_eof537: cs = 537; goto _test_eof
	_test_eof538: cs = 538; goto _test_eof
	_test_eof539: cs = 539; goto _test_eof
	_test_eof540: cs = 540; goto _test_eof
	_test_eof541: cs = 541; goto _test_eof
	_test_eof542: cs = 542; goto _test_eof
	_test_eof543: cs = 543; goto _test_eof
	_test_eof544: cs = 544; goto _test_eof
	_test_eof545: cs = 545; goto _test_eof
	_test_eof546: cs = 546; goto _test_eof
	_test_eof547: cs = 547; goto _test_eof
	_test_eof548: cs = 548; goto _test_eof
	_test_eof549: cs = 549; goto _test_eof
	_test_eof550: cs = 550; goto _test_eof
	_test_eof551: cs = 551; goto _test_eof
	_test_eof552: cs = 552; goto _test_eof
	_test_eof553: cs = 553; goto _test_eof
	_test_eof554: cs = 554; goto _test_eof
	_test_eof555: cs = 555; goto _test_eof
	_test_eof556: cs = 556; goto _test_eof
	_test_eof557: cs = 557; goto _test_eof
	_test_eof558: cs = 558; goto _test_eof
	_test_eof559: cs = 559; goto _test_eof
	_test_eof560: cs = 560; goto _test_eof
	_test_eof561: cs = 561; goto _test_eof
	_test_eof562: cs = 562; goto _test_eof
	_test_eof563: cs = 563; goto _test_eof
	_test_eof564: cs = 564; goto _test_eof
	_test_eof565: cs = 565; goto _test_eof
	_test_eof566: cs = 566; goto _test_eof
	_test_eof567: cs = 567; goto _test_eof
	_test_eof568: cs = 568; goto _test_eof
	_test_eof569: cs = 569; goto _test_eof
	_test_eof570: cs = 570; goto _test_eof
	_test_eof571: cs = 571; goto _test_eof
	_test_eof572: cs = 572; goto _test_eof
	_test_eof573: cs = 573; goto _test_eof
	_test_eof574: cs = 574; goto _test_eof
	_test_eof575: cs = 575; goto _test_eof
	_test_eof576: cs = 576; goto _test_eof
	_test_eof577: cs = 577; goto _test_eof
	_test_eof578: cs = 578; goto _test_eof
	_test_eof579: cs = 579; goto _test_eof
	_test_eof580: cs = 580; goto _test_eof
	_test_eof581: cs = 581; goto _test_eof
	_test_eof582: cs = 582; goto _test_eof
	_test_eof583: cs = 583; goto _test_eof
	_test_eof584: cs = 584; goto _test_eof
	_test_eof585: cs = 585; goto _test_eof
	_test_eof586: cs = 586; goto _test_eof
	_test_eof587: cs = 587; goto _test_eof
	_test_eof588: cs = 588; goto _test_eof
	_test_eof589: cs = 589; goto _test_eof
	_test_eof590: cs = 590; goto _test_eof
	_test_eof591: cs = 591; goto _test_eof
	_test_eof592: cs = 592; goto _test_eof
	_test_eof593: cs = 593; goto _test_eof
	_test_eof594: cs = 594; goto _test_eof
	_test_eof595: cs = 595; goto _test_eof
	_test_eof596: cs = 596; goto _test_eof
	_test_eof597: cs = 597; goto _test_eof
	_test_eof598: cs = 598; goto _test_eof
	_test_eof599: cs = 599; goto _test_eof
	_test_eof600: cs = 600; goto _test_eof
	_test_eof601: cs = 601; goto _test_eof
	_test_eof602: cs = 602; goto _test_eof
	_test_eof603: cs = 603; goto _test_eof
	_test_eof604: cs = 604; goto _test_eof
	_test_eof605: cs = 605; goto _test_eof
	_test_eof606: cs = 606; goto _test_eof
	_test_eof607: cs = 607; goto _test_eof
	_test_eof608: cs = 608; goto _test_eof
	_test_eof609: cs = 609; goto _test_eof
	_test_eof610: cs = 610; goto _test_eof
	_test_eof611: cs = 611; goto _test_eof
	_test_eof612: cs = 612; goto _test_eof
	_test_eof613: cs = 613; goto _test_eof
	_test_eof614: cs = 614; goto _test_eof
	_test_eof615: cs = 615; goto _test_eof
	_test_eof616: cs = 616; goto _test_eof
	_test_eof617: cs = 617; goto _test_eof
	_test_eof618: cs = 618; goto _test_eof
	_test_eof623: cs = 623; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 6, 8, 10, 12, 14, 16:
//line rfc5424/machine.rl:49

    err = fmt.Errorf("error parsing <nilvalue>");

//line rfc5424/parser.go:13584
		}
	}

	_out: {}
	}

//line rfc5424/parser.rl:38


    if cs < rfc5424_first_final {
      return nil, err
    }

    return &SyslogMessage{
      Header: Header{
        Pri: Pri{
          Prival: *prival,
        },
        Version: *version,
        Timestamp: timestamp,
        Hostname: hostname,
        Appname: appname,
        ProcID: procid,
        MsgID: msgid,
      },
    }, nil
}
