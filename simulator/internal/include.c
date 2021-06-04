// Go's CGO build system is very primitive (to put it politely). It can include
// headers from any location, but can only compile sources in the same directory
// as the Go code. Thus to allow us to use the Mircosoft code as a submodule, we
// have to textually include all of the sources into this file.

#define _CRYPT_HASH_C_
#define _X509_SPT_

// Google sources
#include "Clock.not_c"
#include "Entropy.not_c"
#include "NVMem.not_c"
#include "Run.not_c"

// Most of the sources can be included in any order. However, this file has to
// be included first as it instantiates all of the libraries global variables.
#include "support/Global.not_c"

#include "X509/TpmASN1.c"
#include "X509/X509_ECC.c"
#include "X509/X509_RSA.c"
#include "X509/X509_spt.c"
#include "command/Asymmetric/ECC_Parameters.c"
#include "command/Asymmetric/ECDH_KeyGen.c"
#include "command/Asymmetric/ECDH_ZGen.c"
#include "command/Asymmetric/EC_Ephemeral.c"
#include "command/Asymmetric/RSA_Decrypt.c"
#include "command/Asymmetric/RSA_Encrypt.c"
#include "command/Asymmetric/ZGen_2Phase.c"
#include "command/AttachedComponent/AC_GetCapability.c"
#include "command/AttachedComponent/AC_Send.c"
#include "command/AttachedComponent/AC_spt.c"
#include "command/AttachedComponent/Policy_AC_SendSelect.c"
#include "command/Attestation/Attest_spt.c"
#include "command/Attestation/Certify.c"
#include "command/Attestation/CertifyCreation.c"
#include "command/Attestation/CertifyX509.c"
#include "command/Attestation/GetCommandAuditDigest.c"
#include "command/Attestation/GetSessionAuditDigest.c"
#include "command/Attestation/GetTime.c"
#include "command/Attestation/Quote.c"
#include "command/Capability/GetCapability.c"
#include "command/Capability/TestParms.c"
#include "command/ClockTimer/ClockRateAdjust.c"
#include "command/ClockTimer/ClockSet.c"
#include "command/ClockTimer/ReadClock.c"
#include "command/CommandAudit/SetCommandCodeAuditStatus.c"
#include "command/Context/ContextLoad.c"
#include "command/Context/ContextSave.c"
#include "command/Context/Context_spt.c"
#include "command/Context/EvictControl.c"
#include "command/Context/FlushContext.c"
#include "command/DA/DictionaryAttackLockReset.c"
#include "command/DA/DictionaryAttackParameters.c"
#include "command/Duplication/Duplicate.c"
#include "command/Duplication/Import.c"
#include "command/Duplication/Rewrap.c"
#include "command/EA/PolicyAuthValue.c"
#include "command/EA/PolicyAuthorize.c"
#include "command/EA/PolicyAuthorizeNV.c"
#include "command/EA/PolicyCommandCode.c"
#include "command/EA/PolicyCounterTimer.c"
#include "command/EA/PolicyCpHash.c"
#include "command/EA/PolicyDuplicationSelect.c"
#include "command/EA/PolicyGetDigest.c"
#include "command/EA/PolicyLocality.c"
#include "command/EA/PolicyNV.c"
#include "command/EA/PolicyNameHash.c"
#include "command/EA/PolicyNvWritten.c"
#include "command/EA/PolicyOR.c"
#include "command/EA/PolicyPCR.c"
#include "command/EA/PolicyPassword.c"
#include "command/EA/PolicyPhysicalPresence.c"
#include "command/EA/PolicySecret.c"
#include "command/EA/PolicySigned.c"
#include "command/EA/PolicyTemplate.c"
#include "command/EA/PolicyTicket.c"
#include "command/EA/Policy_spt.c"
#include "command/Ecdaa/Commit.c"
#include "command/FieldUpgrade/FieldUpgradeData.c"
#include "command/FieldUpgrade/FieldUpgradeStart.c"
#include "command/FieldUpgrade/FirmwareRead.c"
#include "command/HashHMAC/EventSequenceComplete.c"
#include "command/HashHMAC/HMAC_Start.c"
#include "command/HashHMAC/HashSequenceStart.c"
#include "command/HashHMAC/MAC_Start.c"
#include "command/HashHMAC/SequenceComplete.c"
#include "command/HashHMAC/SequenceUpdate.c"
#include "command/Hierarchy/ChangeEPS.c"
#include "command/Hierarchy/ChangePPS.c"
#include "command/Hierarchy/Clear.c"
#include "command/Hierarchy/ClearControl.c"
#include "command/Hierarchy/CreatePrimary.c"
#include "command/Hierarchy/HierarchyChangeAuth.c"
#include "command/Hierarchy/HierarchyControl.c"
#include "command/Hierarchy/SetPrimaryPolicy.c"
#include "command/Misc/PP_Commands.c"
#include "command/Misc/SetAlgorithmSet.c"
#include "command/NVStorage/NV_Certify.c"
#include "command/NVStorage/NV_ChangeAuth.c"
#include "command/NVStorage/NV_DefineSpace.c"
#include "command/NVStorage/NV_Extend.c"
#include "command/NVStorage/NV_GlobalWriteLock.c"
#include "command/NVStorage/NV_Increment.c"
#include "command/NVStorage/NV_Read.c"
#include "command/NVStorage/NV_ReadLock.c"
#include "command/NVStorage/NV_ReadPublic.c"
#include "command/NVStorage/NV_SetBits.c"
#include "command/NVStorage/NV_UndefineSpace.c"
#include "command/NVStorage/NV_UndefineSpaceSpecial.c"
#include "command/NVStorage/NV_Write.c"
#include "command/NVStorage/NV_WriteLock.c"
#include "command/NVStorage/NV_spt.c"
#include "command/Object/ActivateCredential.c"
#include "command/Object/Create.c"
#include "command/Object/CreateLoaded.c"
#include "command/Object/Load.c"
#include "command/Object/LoadExternal.c"
#include "command/Object/MakeCredential.c"
#include "command/Object/ObjectChangeAuth.c"
#include "command/Object/Object_spt.c"
#include "command/Object/ReadPublic.c"
#include "command/Object/Unseal.c"
#include "command/PCR/PCR_Allocate.c"
#include "command/PCR/PCR_Event.c"
#include "command/PCR/PCR_Extend.c"
#include "command/PCR/PCR_Read.c"
#include "command/PCR/PCR_Reset.c"
#include "command/PCR/PCR_SetAuthPolicy.c"
#include "command/PCR/PCR_SetAuthValue.c"
#include "command/Random/GetRandom.c"
#include "command/Random/StirRandom.c"
#include "command/Session/PolicyRestart.c"
#include "command/Session/StartAuthSession.c"
#include "command/Signature/Sign.c"
#include "command/Signature/VerifySignature.c"
#include "command/Startup/Shutdown.c"
#include "command/Startup/Startup.c"
#include "command/Symmetric/EncryptDecrypt.c"
#include "command/Symmetric/EncryptDecrypt2.c"
#include "command/Symmetric/EncryptDecrypt_spt.c"
#include "command/Symmetric/HMAC.c"
#include "command/Symmetric/Hash.c"
#include "command/Symmetric/MAC.c"
#include "command/Testing/GetTestResult.c"
#include "command/Testing/IncrementalSelfTest.c"
#include "command/Testing/SelfTest.c"
#include "command/Vendor/Vendor_TCG_Test.c"
#include "crypt/AlgorithmTests.c"
#include "crypt/BnConvert.c"
#include "crypt/BnMath.c"
#include "crypt/BnMemory.c"
#include "crypt/CryptCmac.c"
#include "crypt/CryptDes.c"
#include "crypt/CryptEccData.c"
#include "crypt/CryptEccKeyExchange.c"
#include "crypt/CryptEccMain.c"
#include "crypt/CryptEccSignature.c"
#include "crypt/CryptHash.c"
#include "crypt/CryptPrime.c"
#include "crypt/CryptPrimeSieve.c"
#include "crypt/CryptRand.c"
#include "crypt/CryptRsa.c"
#include "crypt/CryptSelfTest.c"
#include "crypt/CryptSmac.c"
#include "crypt/CryptSym.c"
#include "crypt/CryptUtil.c"
#include "crypt/PrimeData.c"
#include "crypt/RsaKeyCache.c"
#include "crypt/Ticket.c"
#include "crypt/ossl/TpmToOsslDesSupport.c"
#include "crypt/ossl/TpmToOsslMath.c"
#include "crypt/ossl/TpmToOsslSupport.c"
#include "events/_TPM_Hash_Data.c"
#include "events/_TPM_Hash_End.c"
#include "events/_TPM_Hash_Start.c"
#include "events/_TPM_Init.c"
#include "main/CommandDispatcher.c"
#include "main/ExecCommand.c"
#include "main/SessionProcess.c"
#include "subsystem/CommandAudit.c"
#include "subsystem/DA.c"
#include "subsystem/Hierarchy.c"
#include "subsystem/NvDynamic.c"
#include "subsystem/NvReserved.c"
#include "subsystem/Object.c"
#include "subsystem/PCR.c"
#include "subsystem/PP.c"
#include "subsystem/Session.c"
#include "subsystem/Time.c"
#include "support/AlgorithmCap.not_c"
#include "support/Bits.not_c"
#include "support/CommandCodeAttributes.not_c"
#include "support/Entity.not_c"
#include "support/Handle.not_c"
#include "support/IoBuffers.not_c"
#include "support/Locality.not_c"
#include "support/Manufacture.not_c"
#include "support/Marshal.not_c"
#include "support/MathOnByteBuffers.not_c"
#include "support/Memory.not_c"
#include "support/Power.not_c"
#include "support/PropertyCap.not_c"
#include "support/Response.not_c"
#include "support/ResponseCodeProcessing.not_c"
#include "support/TpmFail.not_c"
#include "support/TpmSizeChecks.not_c"
