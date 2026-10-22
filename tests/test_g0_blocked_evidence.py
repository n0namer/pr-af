"""Read-only checks of the G0 blocker record; never proof of live DEV identity.

Run without product dependencies: py -m unittest discover -s tests
    -p test_g0_blocked_evidence.py -v
"""

import subprocess
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
G0_HEADER = "### G0 — CURRENT implementation/runtime identity"
IDENTITY_FIELDS = (
    "Python vs Go implementation",
    "exact source/worktree",
    "exact process/artifact",
    "health",
    "registration/reasoner surface and real installed count",
    "callback/public URL routing",
)


def blocked_record_errors(plan):
    """Validate this attempt's documentation, not a runtime discovery response."""
    if not isinstance(plan, str) or not plan.strip():
        return ["missing plan"]
    if plan.count(G0_HEADER) != 1:
        return ["missing or duplicate G0 gate"]
    gate = plan.split(G0_HEADER, 1)[1].split("### ", 1)[0]
    errors = []
    if not gate.lstrip().startswith("Status: BLOCKED"):
        errors.append("G0 must remain explicitly BLOCKED")
    for field in IDENTITY_FIELDS:
        if f"- {field}: UNAVAILABLE / UNVERIFIED_CURRENT" not in gate:
            errors.append(f"missing unavailable identity: {field}")
    for text in (
        "no AgentField DEV/runtime connector",
        "AF_COMMAND=UNAVAILABLE",
        "This offline package test does not establish any CURRENT DEV runtime identity.",
        "No AgentField DEV, production, callback, or source mutation was attempted.",
        "An authorized AgentField DEV read-only interface/operator",
    ):
        if text not in gate:
            errors.append(f"missing evidence or boundary: {text}")
    return errors


class TestG0BlockedEvidence(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        cls.plan = (ROOT / "PLAN.md").read_text(encoding="utf-8")

    def test_ac1_current_identity_explicitly_blocked(self):
        self.assertEqual(blocked_record_errors(self.plan), [])

    def test_committed_plan_has_blocked_evidence_without_worktree_additions(self):
        result = subprocess.run(
            ["git", "-c", f"safe.directory={ROOT.as_posix()}", "show",
             "ee7467243a39891595a8c5e61812dbb16864256b:PLAN.md"],
            cwd=ROOT, capture_output=True, text=True, encoding="utf-8", check=True,
        )
        self.assertEqual(blocked_record_errors(result.stdout), [])

    def test_ac2_installed_surface_and_process_explicitly_unavailable(self):
        for field in (IDENTITY_FIELDS[2], IDENTITY_FIELDS[4]):
            with self.subTest(field=field):
                self.assertIn(f"- {field}: UNAVAILABLE / UNVERIFIED_CURRENT", self.plan)

    def test_ac3_candidate_commit_has_no_source_or_callback_changes(self):
        result = subprocess.run(
            ["git", "-c", f"safe.directory={ROOT.as_posix()}", "diff-tree",
             "--no-commit-id", "--name-only", "-r", "ee7467243a39891595a8c5e61812dbb16864256b"],
            cwd=ROOT, capture_output=True, text=True, check=True,
        )
        self.assertEqual(result.stdout.splitlines(), ["PLAN.md"])

    def test_ac4_single_canonical_gate_and_no_runtime_canary_claim(self):
        self.assertEqual(self.plan.count(G0_HEADER), 1)
        self.assertIn("Source of truth: this file (`PLAN.md`)", self.plan)
        self.assertIn("Evidence/tests/canaries: No runtime readback or canary.", self.plan)
        self.assertIn("Status: BLOCKED_BY_G0", self.plan)

    def test_handoff_sha_is_historical_documentation_not_runtime_evidence(self):
        handoff = self.plan.split("### G0 attempt handoff — 2026-10-20", 1)[1].split("### ", 1)[0]
        self.assertNotIn("- Durable commit/SHA if any: none.", handoff)
        self.assertIn("none at the time this entry was written", handoff)
        self.assertIn(
            "documentation-only handoff was later committed as "
            "`ee7467243a39891595a8c5e61812dbb16864256b`", handoff,
        )
        self.assertIn(
            "No accepted runtime delta or runtime identity evidence is associated with that commit.", handoff,
        )

    def test_blocked_record_accepts_windows_line_endings_and_final_gate_boundary(self):
        # Following gates must not be needed to validate the G0 blocker record.
        gate_only = G0_HEADER + self.plan.split(G0_HEADER, 1)[1].split("### ", 1)[0]
        for plan in (gate_only.rstrip(), self.plan.replace("\n", "\r\n")):
            with self.subTest(windows_line_endings="\r\n" in plan):
                self.assertEqual(blocked_record_errors(plan), [])

    def test_missing_none_empty_and_whitespace_plan_rejected(self):
        for plan in (None, "", " \n\t", 0):
            with self.subTest(plan=plan):
                self.assertEqual(blocked_record_errors(plan), ["missing plan"])

    def test_missing_and_duplicate_gate_rejected(self):
        for plan in ("# No G0 evidence", self.plan + "\n" + G0_HEADER):
            with self.subTest(plan_length=len(plan)):
                self.assertEqual(blocked_record_errors(plan), ["missing or duplicate G0 gate"])

    def test_each_missing_identity_field_rejected(self):
        for field in IDENTITY_FIELDS:
            with self.subTest(field=field):
                altered = self.plan.replace(f"- {field}: UNAVAILABLE / UNVERIFIED_CURRENT", "")
                self.assertIn(f"missing unavailable identity: {field}", blocked_record_errors(altered))

    def test_zero_or_expected_count_cannot_substitute_for_installed_evidence(self):
        field = IDENTITY_FIELDS[4]
        for count in ("0", "17", "None", ""):
            with self.subTest(count=count):
                altered = self.plan.replace(
                    f"- {field}: UNAVAILABLE / UNVERIFIED_CURRENT", f"- {field}: {count}")
                self.assertTrue(blocked_record_errors(altered))

    def test_pass_or_empty_status_without_runtime_evidence_rejected(self):
        for status in ("PASS", "", "UNAVAILABLE"):
            with self.subTest(status=status):
                altered = self.plan.replace("Status: BLOCKED —", f"Status: {status} —", 1)
                self.assertIn("G0 must remain explicitly BLOCKED", blocked_record_errors(altered))

    def test_offline_success_disclaimer_and_mutation_boundary_required(self):
        for text in (
            "This offline package test does not establish any CURRENT DEV runtime identity.",
            "No AgentField DEV, production, callback, or source mutation was attempted.",
        ):
            with self.subTest(text=text):
                self.assertTrue(blocked_record_errors(self.plan.replace(text, "")))

    def test_missing_access_evidence_and_recovery_path_rejected(self):
        for text in (
            "no AgentField DEV/runtime connector",
            "AF_COMMAND=UNAVAILABLE",
            "An authorized AgentField DEV read-only interface/operator",
        ):
            with self.subTest(text=text):
                self.assertIn(
                    f"missing evidence or boundary: {text}",
                    blocked_record_errors(self.plan.replace(text, "")),
                )

    def test_identity_outside_gate_cannot_supply_missing_g0_evidence(self):
        field = IDENTITY_FIELDS[5]
        evidence = f"- {field}: UNAVAILABLE / UNVERIFIED_CURRENT"
        for outside in (f"{evidence}\n", f"\n### Other gate\n{evidence}"):
            with self.subTest(outside=outside):
                altered = self.plan.replace(evidence, "")
                if outside.startswith("\n"):
                    altered += outside
                else:
                    altered = outside + altered
                self.assertIn(
                    f"missing unavailable identity: {field}", blocked_record_errors(altered)
                )


if __name__ == "__main__":
    unittest.main()
